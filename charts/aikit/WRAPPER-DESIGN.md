# AIKit Distributed Inference Architecture Design

This document outlines the wrapper components needed to enable distributed inference with LeaderWorkerSet for AIKit.

## Current Architecture

AIKit currently builds container images with:
- **LocalAI binary** (`/usr/bin/local-ai`) as the entrypoint
- **Model files** and **configuration** embedded in the image
- **Single-node inference** using LocalAI directly

## Distributed Architecture Requirements

For LeaderWorkerSet distributed inference, we need coordination components similar to the [llamacpp LWS example](https://github.com/kubernetes-sigs/lws/tree/main/docs/examples/llamacpp).

### Required Components

#### 1. Leader Wrapper (`aikit-leader`)

**Purpose**: Coordinates distributed inference by discovering workers and configuring LocalAI for multi-node operation.

**Key Responsibilities**:
- Discover worker pods using LWS environment variables
- Wait for all workers to be ready
- Configure LocalAI for distributed inference mode
- Start LocalAI with appropriate distributed settings

**Implementation** (similar to `llamacpp-leader/main.go`):

```go
package main

import (
    "context"
    "fmt"
    "net"
    "os"
    "os/exec"
    "strconv"
    "strings"
    "time"
)

func main() {
    if err := run(context.Background()); err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
        os.Exit(1)
    }
}

func run(ctx context.Context) error {
    // Get LWS configuration from environment
    lwsSize := getLWSGroupSize()
    lwsLeaderAddress := os.Getenv("LWS_LEADER_ADDRESS")

    // Discover worker pods
    workers := discoverWorkers(lwsSize, lwsLeaderAddress)

    // Configure LocalAI for distributed mode
    configureDistributedLocalAI(workers)

    // Start LocalAI with distributed configuration
    return startLocalAI(ctx)
}

func getLWSGroupSize() int {
    if s := os.Getenv("LWS_GROUP_SIZE"); s != "" {
        if v, err := strconv.Atoi(s); err == nil {
            return v
        }
    }
    return 1 // fallback to single node
}

func discoverWorkers(lwsSize int, lwsLeaderAddress string) []string {
    serviceTokens := strings.Split(lwsLeaderAddress, ".")
    var workers []string

    // Workers are indexed 1, 2, 3... (0 is leader)
    for i := 1; i < lwsSize; i++ {
        workerHost := fmt.Sprintf("%s-%d.%s",
            serviceTokens[0], i, strings.Join(serviceTokens[1:], "."))

        // Wait for worker to be resolvable
        if ip := waitForWorker(workerHost); ip != "" {
            workers = append(workers, ip)
        }
    }
    return workers
}

func waitForWorker(host string) string {
    maxAttempts := 30
    for attempt := 0; attempt < maxAttempts; attempt++ {
        if ips, err := net.LookupIP(host); err == nil && len(ips) > 0 {
            return ips[0].String()
        }
        time.Sleep(5 * time.Second)
    }
    return ""
}

func configureDistributedLocalAI(workers []string) error {
    // Generate LocalAI configuration for distributed inference
    // This depends on LocalAI's distributed capabilities

    // Example: Create a config file that tells LocalAI about workers
    config := generateDistributedConfig(workers)
    return writeConfigFile("/config.yaml", config)
}

func startLocalAI(ctx context.Context) error {
    args := []string{
        "--config-file=/config.yaml",
        "--address=0.0.0.0",
        "--port=8080",
    }

    cmd := exec.CommandContext(ctx, "/usr/bin/local-ai", args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    return cmd.Run()
}
```

#### 2. Worker Wrapper (`aikit-worker`)

**Purpose**: Connects to the leader and provides distributed compute capacity.

**Key Responsibilities**:
- Connect to leader coordinator
- Run LocalAI in worker mode or run specialized worker processes
- Handle worker-specific configuration

**Implementation Options**:

**Option A**: LocalAI Worker Mode (if LocalAI supports it)
```go
func main() {
    leaderAddress := os.Getenv("LEADER_ADDRESS")

    args := []string{
        "--worker-mode",
        "--leader=" + leaderAddress,
        "--worker-port=50051",
    }

    cmd := exec.Command("/usr/bin/local-ai", args...)
    cmd.Run()
}
```

**Option B**: Specialized Worker Process (if LocalAI doesn't support worker mode)
```go
func main() {
    // Run a gRPC/HTTP server that the leader can coordinate with
    // This would require implementing the worker protocol
    startWorkerServer()
}
```

### LocalAI Distributed Inference Support

**Current Challenge**: LocalAI may not have built-in distributed inference support like llama.cpp's RPC mode.

**Solutions**:

1. **Check LocalAI Documentation**: Investigate if LocalAI has distributed/cluster modes
2. **Implement Custom Coordination**: Create coordination layer that:
   - Distributes requests across workers
   - Handles model sharding/partitioning
   - Manages inter-worker communication

3. **Alternative Backends**: Consider other inference engines with better distributed support:
   - vLLM (has tensor parallel support)
   - TensorRT-LLM (supports multi-node)
   - Text Generation Inference (TGI)

### Updated Helm Chart Integration

The wrapper components would be used in the LeaderWorkerSet template like this:

```yaml
leaderWorkerSet:
  enabled: true
  leader:
    image:
      repository: "ghcr.io/sozercan/aikit-leader"
      tag: "latest"
    command: ["/aikit-leader"]
    args:
      - "--model-path=/models"
      - "--config=/config.yaml"
    env:
      - name: LWS_GROUP_SIZE
        valueFrom:
          fieldRef:
            fieldPath: metadata.annotations['leaderworkerset.sigs.k8s.io/group-size']
      - name: LWS_LEADER_ADDRESS
        valueFrom:
          fieldRef:
            fieldPath: metadata.annotations['leaderworkerset.sigs.k8s.io/leader-address']

  worker:
    image:
      repository: "ghcr.io/sozercan/aikit-worker"
      tag: "latest"
    command: ["/aikit-worker"]
    args:
      - "--leader-address=$(LWS_LEADER_ADDRESS)"
```

## Implementation Steps

### Phase 1: Research & Design
1. **Investigate LocalAI distributed capabilities**
2. **Design coordination protocol** between leader/workers
3. **Define worker communication interfaces**

### Phase 2: Core Components
1. **Implement `aikit-leader`** wrapper
2. **Implement `aikit-worker`** wrapper (if needed)
3. **Create distributed LocalAI configuration generation**

### Phase 3: Container Images
1. **Build leader container image** with wrapper + LocalAI
2. **Build worker container image** with wrapper + LocalAI
3. **Update AIKit build process** to support wrapper images

### Phase 4: Integration & Testing
1. **Update Helm templates** to use wrapper images
2. **Test distributed inference** with multiple models
3. **Performance optimization** and scaling tests

## Alternative Approaches

### Option 1: LocalAI Plugin/Extension
- Extend LocalAI with distributed inference plugins
- Minimal wrapper code needed
- Requires LocalAI core changes

### Option 2: Proxy/Gateway Pattern
- AIKit-specific distributed gateway
- Routes requests to LocalAI instances
- More control over load balancing and coordination

### Option 3: Alternative Inference Engine
- Replace LocalAI with natively distributed engine
- Requires significant AIKit architecture changes
- Better long-term distributed performance

## Recommended Next Steps

1. **Research LocalAI distributed capabilities** - check if newer versions support clustering
2. **Prototype simple leader wrapper** - basic coordination without full distributed inference
3. **Test with existing models** - validate the wrapper approach works
4. **Consider inference engine alternatives** - evaluate vLLM, TGI, or TensorRT-LLM integration

The wrapper approach provides a path to distributed inference while maintaining compatibility with the existing AIKit architecture.

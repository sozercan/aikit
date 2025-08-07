# AIKit Distributed Inference with LeaderWorkerSet

This directory contains Helm chart configurations for deploying AIKit with distributed inference capabilities using Kubernetes [LeaderWorkerSet (LWS)](https://github.com/kubernetes-sigs/lws).

## Overview

LeaderWorkerSet enables distributed inference for large language models by orchestrating a leader pod (which handles requests and coordination) and multiple worker pods (which provide additional compute resources). This is particularly useful for:

- Large models that benefit from tensor parallelism
- Distributed compute across multiple nodes
- Improved resource utilization and scalability

## Prerequisites

1. **Kubernetes cluster** with GPU support (for GPU-accelerated inference)
2. **LeaderWorkerSet controller** installed in your cluster
3. **NVIDIA GPU Operator** (if using GPU acceleration)

### Installing LeaderWorkerSet

```bash
# Install LeaderWorkerSet CRDs and controller
kubectl apply --server-side -f https://github.com/kubernetes-sigs/lws/releases/download/v0.4.0/manifests.yaml
```

## Configuration

### Standard Deployment (Default)

The default configuration deploys AIKit as a standard Kubernetes Deployment:

```bash
helm install aikit ./charts/aikit
```

### Distributed Inference with LeaderWorkerSet

To enable distributed inference, set `leaderWorkerSet.enabled: true`:

```bash
helm install aikit ./charts/aikit --set leaderWorkerSet.enabled=true
```

Or use the provided distributed configuration:

```bash
helm install aikit ./charts/aikit -f charts/aikit/values-distributed.yaml
```

### Key Configuration Parameters

#### LeaderWorkerSet Configuration

- `leaderWorkerSet.enabled`: Enable/disable LeaderWorkerSet mode
- `leaderWorkerSet.replicas`: Number of LeaderWorkerSet replicas
- `leaderWorkerSet.leaderWorkerTemplate.size`: Total pods per replica (1 leader + N workers)
- `leaderWorkerSet.leaderWorkerTemplate.restartPolicy`: Pod restart behavior

#### Leader Pod Configuration

```yaml
leaderWorkerSet:
  leader:
    image:
      repository: ""  # Uses main image if empty
      tag: ""         # Uses main image tag if empty
      pullPolicy: ""  # Uses main image pullPolicy if empty
    resources:
      limits:
        memory: 32Gi
        nvidia.com/gpu: "2"
    args:
      - "--host"
      - "0.0.0.0"
      - "--n-gpu-layers"
      - "99"
    env:
      - name: LWS_GROUP_SIZE
        value: "3"
```

#### Worker Pod Configuration

```yaml
leaderWorkerSet:
  worker:
    image:
      repository: "aikit-worker"
      tag: "latest"
    resources:
      limits:
        memory: 16Gi
        nvidia.com/gpu: "2"
    args:
      - "--host"
      - "0.0.0.0"
      - "--mem"
      - "8192"
```

## Example Configurations

### CPU-based Distributed Inference

For testing or CPU-only environments:

```yaml
leaderWorkerSet:
  enabled: true
  leaderWorkerTemplate:
    size: 3  # 1 leader + 2 workers
  leader:
    resources:
      limits:
        memory: 8Gi
      requests:
        cpu: 2000m
        memory: 4Gi
  worker:
    resources:
      limits:
        memory: 4Gi
      requests:
        cpu: 1000m
        memory: 2Gi
```

### GPU-accelerated Distributed Inference

For production workloads with GPU acceleration:

```yaml
leaderWorkerSet:
  enabled: true
  leaderWorkerTemplate:
    size: 5  # 1 leader + 4 workers
  leader:
    resources:
      limits:
        nvidia.com/gpu: "4"
        memory: 64Gi
  worker:
    resources:
      limits:
        nvidia.com/gpu: "2"
        memory: 32Gi

nodeSelector:
  node.kubernetes.io/instance-type: "g5.4xlarge"

tolerations:
  - key: nvidia.com/gpu
    operator: Exists
    effect: NoSchedule
```

## Usage

### Deploying with LeaderWorkerSet

1. **Basic deployment:**
   ```bash
   helm install aikit ./charts/aikit \
     --set leaderWorkerSet.enabled=true \
     --set leaderWorkerSet.leaderWorkerTemplate.size=3
   ```

2. **Using the distributed configuration:**
   ```bash
   helm install aikit ./charts/aikit -f charts/aikit/values-distributed.yaml
   ```

3. **Custom configuration:**
   ```bash
   helm install aikit ./charts/aikit \
     --set leaderWorkerSet.enabled=true \
     --set leaderWorkerSet.replicas=2 \
     --set leaderWorkerSet.leaderWorkerTemplate.size=4 \
     --set image.repository=ghcr.io/sozercan/llama3.3 \
     --set image.tag=70b
   ```

### Accessing the Service

The service works the same way regardless of deployment mode:

```bash
# Port forward to access the service
kubectl port-forward svc/aikit 8080:8080

# Test the API
curl http://localhost:8080/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "llama-3.3-70b-instruct",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
```

### Monitoring

Check the status of your LeaderWorkerSet deployment:

```bash
# Check LeaderWorkerSet status
kubectl get leaderworkerset

# Check individual pods
kubectl get pods -l app.kubernetes.io/name=aikit

# Check logs from leader pod
kubectl logs -l app.kubernetes.io/name=aikit,leaderworkerset.sigs.k8s.io/pod-role=leader

# Check logs from worker pods
kubectl logs -l app.kubernetes.io/name=aikit,leaderworkerset.sigs.k8s.io/pod-role=worker
```

## Limitations

- **Horizontal Pod Autoscaling (HPA)** is not supported with LeaderWorkerSet
- **Model sharing** between leader and workers depends on the specific model implementation
- **Node affinity** may be required for optimal performance in multi-node setups

## Advanced Configuration

### Multi-Node Deployment

For optimal performance across multiple nodes:

```yaml
affinity:
  podAntiAffinity:
    preferredDuringSchedulingIgnoredDuringExecution:
    - weight: 100
      podAffinityTerm:
        labelSelector:
          matchExpressions:
          - key: app.kubernetes.io/name
            operator: In
            values:
            - aikit
        topologyKey: kubernetes.io/hostname
```

### Custom Worker Images

You can specify different images for leader and worker pods:

```yaml
leaderWorkerSet:
  leader:
    image:
      repository: "ghcr.io/sozercan/llama3.3"
      tag: "70b"
  worker:
    image:
      repository: "custom/aikit-worker"
      tag: "latest"
```

## Troubleshooting

### Common Issues

1. **Pods stuck in Pending state:**
   - Check resource requests vs. available cluster resources
   - Verify GPU availability if using GPU resources
   - Check node selectors and tolerations

2. **Leader can't communicate with workers:**
   - Ensure network policies allow pod-to-pod communication
   - Check firewall rules and security groups

3. **Performance issues:**
   - Consider node placement and network bandwidth between nodes
   - Adjust resource allocations
   - Review model-specific configuration

### Debugging Commands

```bash
# Check LeaderWorkerSet events
kubectl describe leaderworkerset aikit

# Check pod events
kubectl describe pod <pod-name>

# Check resource utilization
kubectl top pod -l app.kubernetes.io/name=aikit

# Check network connectivity between pods
kubectl exec -it <leader-pod> -- ping <worker-pod-ip>
```

## Examples from LWS Community

The LeaderWorkerSet community provides several examples that can be adapted for AIKit:

- [llamacpp example](https://github.com/kubernetes-sigs/lws/tree/main/docs/examples/llamacpp)
- [vLLM example](https://github.com/kubernetes-sigs/lws/tree/main/docs/examples/vllm)
- [TensorRT-LLM example](https://github.com/kubernetes-sigs/lws/tree/main/docs/examples/tensorrt-llm)
- [SGLang example](https://github.com/kubernetes-sigs/lws/tree/main/docs/examples/sglang)

## Contributing

To contribute improvements to the distributed inference configuration:

1. Test your changes with both single-node and multi-node setups
2. Update documentation for any new configuration options
3. Consider backward compatibility with existing deployments
4. Add appropriate validation for LeaderWorkerSet-specific settings

# AIKit Helm Chart

A Kubernetes Helm chart for deploying AIKit with support for both standard deployment and distributed inference using LeaderWorkerSet.

## Overview

AIKit is a comprehensive platform to quickly get started to host, deploy, build and fine-tune large language models (LLMs). This Helm chart supports:

- **Standard Deployment**: Traditional Kubernetes Deployment for single-node inference
- **Distributed Inference**: Multi-node distributed inference using [LeaderWorkerSet](https://github.com/kubernetes-sigs/lws)

## Quick Start

### Standard Deployment

```bash
helm install aikit ./charts/aikit
```

### Distributed Inference

```bash
# Install LeaderWorkerSet first (if not already installed)
kubectl apply --server-side -f https://github.com/kubernetes-sigs/lws/releases/download/v0.4.0/manifests.yaml

# Deploy AIKit with distributed inference
helm install aikit ./charts/aikit --set leaderWorkerSet.enabled=true
```

## Configuration

### Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `image.repository` | string | `"ghcr.io/sozercan/llama3"` | Container image repository |
| `image.tag` | string | `"8b"` | Container image tag |
| `image.pullPolicy` | string | `"IfNotPresent"` | Image pull policy |
| `replicaCount` | int | `1` | Number of replicas (standard mode only) |
| `leaderWorkerSet.enabled` | bool | `false` | Enable distributed inference with LeaderWorkerSet |
| `leaderWorkerSet.replicas` | int | `1` | Number of LeaderWorkerSet replicas |
| `leaderWorkerSet.leaderWorkerTemplate.size` | int | `3` | Total pods per replica (1 leader + N workers) |
| `service.type` | string | `"ClusterIP"` | Kubernetes service type |
| `service.port` | int | `8080` | Service port |
| `resources.limits.memory` | string | `"8Gi"` | Memory limit |
| `resources.requests.cpu` | string | `"100m"` | CPU request |

### Example Configurations

The chart includes several example configurations:

- `values-distributed.yaml`: General distributed inference setup
- `values-llamacpp-style.yaml`: CPU-based distributed inference
- `values-vllm-style.yaml`: GPU-accelerated distributed inference

## Distributed Inference with LeaderWorkerSet

### Prerequisites

1. Kubernetes cluster with [LeaderWorkerSet](https://github.com/kubernetes-sigs/lws) installed
2. GPU support (for GPU-accelerated inference)

### Configuration

Enable distributed inference by setting `leaderWorkerSet.enabled: true`:

```yaml
leaderWorkerSet:
  enabled: true
  replicas: 1
  leaderWorkerTemplate:
    size: 3  # 1 leader + 2 workers
    restartPolicy: RecreateGroupOnPodRestart

  leader:
    resources:
      limits:
        memory: 32Gi
        nvidia.com/gpu: "2"
    args:
      - "--host"
      - "0.0.0.0"
      - "--n-gpu-layers"
      - "99"

  worker:
    image:
      repository: "aikit-worker"
      tag: "latest"
    resources:
      limits:
        memory: 16Gi
        nvidia.com/gpu: "2"
```

### Deployment Examples

#### CPU-based Inference
```bash
helm install aikit ./charts/aikit -f charts/aikit/values-llamacpp-style.yaml
```

#### GPU-accelerated Inference
```bash
helm install aikit ./charts/aikit -f charts/aikit/values-vllm-style.yaml
```

#### Custom Configuration
```bash
helm install aikit ./charts/aikit \
  --set leaderWorkerSet.enabled=true \
  --set leaderWorkerSet.leaderWorkerTemplate.size=5 \
  --set image.repository=ghcr.io/sozercan/llama3.3 \
  --set image.tag=70b
```

## Usage

### Accessing the Service

```bash
# Port forward to access locally
kubectl port-forward svc/aikit 8080:8080

# Access WebUI
open http://localhost:8080/chat

# Test API
curl http://localhost:8080/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "llama-3.1-8b-instruct",
    "messages": [{"role": "user", "content": "Hello!"}]
  }'
```

### Monitoring

```bash
# Check deployment status
kubectl get deployment aikit  # Standard mode
kubectl get leaderworkerset aikit  # Distributed mode

# Check pods
kubectl get pods -l app.kubernetes.io/name=aikit

# View logs
kubectl logs -l app.kubernetes.io/name=aikit
```

## Advanced Configuration

### Multi-Node GPU Setup

```yaml
leaderWorkerSet:
  enabled: true
  leaderWorkerTemplate:
    size: 4

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

nodeSelector:
  node.kubernetes.io/instance-type: "g5.4xlarge"

tolerations:
  - key: nvidia.com/gpu
    operator: Exists
    effect: NoSchedule
```

### Custom Resource Allocation

```yaml
leaderWorkerSet:
  leader:
    resources:
      limits:
        nvidia.com/gpu: "4"
        memory: 64Gi
      requests:
        cpu: 4000m
        memory: 32Gi
  worker:
    resources:
      limits:
        nvidia.com/gpu: "2"
        memory: 32Gi
      requests:
        cpu: 2000m
        memory: 16Gi
```

## Limitations

- Horizontal Pod Autoscaling (HPA) is not supported with LeaderWorkerSet
- Distributed inference requires compatible model implementations
- GPU sharing between leader and workers depends on model architecture

## Documentation

For detailed information about distributed inference:
- [Distributed Inference Guide](./DISTRIBUTED-INFERENCE.md)
- [LeaderWorkerSet Documentation](https://github.com/kubernetes-sigs/lws)
- [AIKit Documentation](https://sozercan.github.io/aikit/)

## Contributing

Contributions are welcome! Please ensure:
1. Test both standard and distributed deployment modes
2. Update documentation for new features
3. Maintain backward compatibility
4. Follow Helm best practices

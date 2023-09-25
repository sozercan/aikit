# AIKit ✨

AIKit is a quick and easy way to get started to host and deploy large language models (LLMs) for inference. No GPU or additional tools needed to get started except for Docker!

AIKit uses [LocalAI](https://localai.io/) under-the-hood to run inference. LocalAI provides a drop-in replacement REST API that is OpenAI API compatible, so you can use any OpenAI API compatible client, such as [Kubectl AI](https://github.com/sozercan/kubectl-ai), to send requests to open-source LLMs powered by AIKit!

## Features
- 🐳 No GPU, internet access or additional tools needed except for Docker!
- 🤏 Minimal image size, vulnerabilities and attack surface with a [distroless](https://github.com/GoogleContainerTools/distroless)-based image
- 🚀 Easy to use declarative configuration
- ✨ OpenAI API compatible to use with any OpenAI API compatible client
- 🚢 Kubernetes deployment ready
- 📦 Supports multiple models in a single image
- 🖥️ Supports GPU-accelerated inferencing with NVIDIA GPUs

## Demo

[![asciicast]()]()

## Quick Start

### Creating a custom model

Create an `aikitfile.yaml` with the following structure:

```yaml
#syntax=docker.io/sozercan/aikit:latest
apiVersion: v1alpha1
models:
  - name: llama-2-7b-chat
    source: https://huggingface.co/TheBloke/Llama-2-7B-Chat-GGUF/resolve/main/llama-2-7b-chat.Q4_K_M.gguf
```

> [!TIP]
> For full `aikitfile` specification, see [specs](docs/specs.md).

> [!TIP]
> See [models folder](./models/) for sample models. [go-skynet/model-gallery](https://github.com/go-skynet/model-gallery) is a good place to find more examples.

Then run the following commands:

```bash
# create a buildx builder
# alternatively, use docker v24 with [containerd image store](https://docs.docker.com/storage/containerd/)) enabled and skip this step
docker buildx create --use --name builder

# build your custom model
docker buildx build -t my-model .
```

> [!TIP]
> If you named your `aikitfile.yaml` something else, you can pass the file name to `docker build` command with `--file` flag.

This will build a docker image with your custom model(s).

```bash
docker images
REPOSITORY    TAG       IMAGE ID       CREATED             SIZE
my-model      latest    e7b7c5a4a2cb   About an hour ago   5.51GB
```

### Running models

You can start the inferencing server for your models with:

```bash
docker run -d --rm -p 8080:8080 my-model
```

You can then send requests to `localhost:8080` to run inference from your models. For example:

```bash
curl http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d '{
     "model": "llama-2-7b-chat",
     "messages": [{"role": "user", "content": "explain kubernetes in a sentence"}]
   }'
{"created":1701236489,"object":"chat.completion","id":"dd1ff40b-31a7-4418-9e32-42151ab6875a","model":"llama-2-7b-chat","choices":[{"index":0,"finish_reason":"stop","message":{"role":"assistant","content":"\nKubernetes is a container orchestration system that automates the deployment, scaling, and management of containerized applications in a microservices architecture."}}],"usage":{"prompt_tokens":0,"completion_tokens":0,"total_tokens":0}}
```

## Kubernetes Deployment

It is easy to get started to deploy your models to Kubernetes!

Make sure you have a Kubernetes cluster running and `kubectl` is configured to talk to it, and your model images are accessible from the cluster. You can use [kind](https://kind.sigs.k8s.io/) to create a local Kubernetes cluster for testing purposes.

```bash
# create a test cluster using kind
kind create cluster

# load your local image to the cluster
kind load docker-image my-model

# create a deployment
kubectl create deployment my-llm-deployment --image=my-model --image-pull-policy=IfNotPresent

# expose it as a service
kubectl expose deployment my-llm-deployment --port=8080 --target-port=8008 --name=my-llm-service

# easy to scale up and down
kubectl scale deployment my-llm-deployment --replicas=3

# port-forward for testing locally
kubectl port-forward service/my-llm-service 8080:8080

# send requests to your model
curl http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d '{
     "model": "llama-2-7b-chat",
     "messages": [{"role": "user", "content": "explain kubernetes in a sentence"}]
   }'
{"created":1701236489,"object":"chat.completion","id":"dd1ff40b-31a7-4418-9e32-42151ab6875a","model":"llama-2-7b-chat","choices":[{"index":0,"finish_reason":"stop","message":{"role":"assistant","content":"\nKubernetes is a container orchestration system that automates the deployment, scaling, and management of containerized applications in a microservices architecture."}}],"usage":{"prompt_tokens":0,"completion_tokens":0,"total_tokens":0}}
```

> [!TIP]
> For an example Kubernetes deployment YAML, see [deployment.yaml](kubernetes/deployment.yaml).

## GPU Support

> [!NOTE]
> At this time, only NVIDIA GPU acceleration is supported.

### NVIDIA

AIKit supports GPU accelerated inferencing with [NVIDIA Container Toolkit](https://github.com/NVIDIA/nvidia-container-toolkit). You must also have the [NVIDIA Drivers](https://www.nvidia.com/en-us/drivers/unix/) installed on your host machine.

For Kubernetes, [NVIDIA GPU Operator](https://github.com/NVIDIA/gpu-operator) provides a streamlined way to install the NVIDIA drivers and container toolkit to configure your cluster to use GPUs.

To get started with GPU-accelerated inferencing, make sure to set the following in your `aikitfile` and build your model.

```yaml
runtime: gpu-nvidia   # use NVIDIA CUDA runtime
f16: true             # use float16 precision
gpu_layers: 35        # number of layers to offload to GPU
low_vram: true        # for devices with low VRAM
```

> [!TIP]
> Make sure to customize these values based in your model and GPU specs.

After building the model, you can run it with [`--gpus all`](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/docker-specialized.html#gpu-enumeration) flag to enable GPU support:

```bash
docker run --rm --gpus all -p 8080:8080 my-model
```

If GPU acceleration is working, you'll see output that is similar to following in the debug logs:

```bash
5:32AM DBG GRPC(llama-2-7b-chat.Q4_K_M.gguf-127.0.0.1:43735): stderr ggml_init_cublas: found 1 CUDA devices:
5:32AM DBG GRPC(llama-2-7b-chat.Q4_K_M.gguf-127.0.0.1:43735): stderr   Device 0: Tesla T4, compute capability 7.5
...
5:32AM DBG GRPC(llama-2-7b-chat.Q4_K_M.gguf-127.0.0.1:43735): stderr llm_load_tensors: using CUDA for GPU acceleration
5:32AM DBG GRPC(llama-2-7b-chat.Q4_K_M.gguf-127.0.0.1:43735): stderr llm_load_tensors: mem required  =   70.41 MB (+ 2048.00 MB per state)
5:32AM DBG GRPC(llama-2-7b-chat.Q4_K_M.gguf-127.0.0.1:43735): stderr llm_load_tensors: offloading 32 repeating layers to GPU
5:32AM DBG GRPC(llama-2-7b-chat.Q4_K_M.gguf-127.0.0.1:43735): stderr llm_load_tensors: offloading non-repeating layers to GPU
5:32AM DBG GRPC(llama-2-7b-chat.Q4_K_M.gguf-127.0.0.1:43735): stderr llm_load_tensors: offloading v cache to GPU
5:32AM DBG GRPC(llama-2-7b-chat.Q4_K_M.gguf-127.0.0.1:43735): stderr llm_load_tensors: offloading k cache to GPU
5:32AM DBG GRPC(llama-2-7b-chat.Q4_K_M.gguf-127.0.0.1:43735): stderr llm_load_tensors: offloaded 35/35 layers to GPU
5:32AM DBG GRPC(llama-2-7b-chat.Q4_K_M.gguf-127.0.0.1:43735): stderr llm_load_tensors: VRAM used: 5869 MB
```

## Acknowledgements

- [LocalAI](https://localai.io/) for providing the inference engine
- [Mockerfile](https://github.com/r2d4/mockerfile) for the inspiration
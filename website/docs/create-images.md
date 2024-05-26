---
title: Creating Model Images
---

:::note
This section shows how to create a custom image with models of your choosing. If you want to use one of the pre-made models, skip to [running models](#running-models).
:::

Create an `aikitfile.yaml` with the following structure:

```yaml
#syntax=ghcr.io/sozercan/aikit:latest
apiVersion: v1alpha1
models:
  - name: llama-2-7b-chat
    source: https://huggingface.co/TheBloke/Llama-2-7B-Chat-GGUF/resolve/main/llama-2-7b-chat.Q4_K_M.gguf
```

:::tip
This is the simplest way to get started to build an image. For full `aikitfile` inference specifications, see [Inference API Specifications](docs/specs-inference.md).
:::

First, create a buildx buildkit instance. Alternatively, if you are using Docker v24 with [containerd image store](https://docs.docker.com/storage/containerd/) enabled, you can skip this step.

```bash
docker buildx create --use --name aikit-builder
```

Then build your image with:

```bash
docker buildx build . -t my-model -f aikitfile.yaml --load
```

This will build a local container image with your model(s). You can see the image with:

```bash
docker images
REPOSITORY    TAG       IMAGE ID       CREATED             SIZE
my-model      latest    e7b7c5a4a2cb   About an hour ago   5.51GB
```

### Running models

You can start the inferencing server for your models with:

```bash
# for pre-made models, replace "my-model" with the image name
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

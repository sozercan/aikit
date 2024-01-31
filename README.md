# AIKit ✨

<p align="center">
<img src="./website/static/img/logo.png" width="100"><br>
</p>

AIKit is a quick, easy, and local or cloud-agnostic way to get started to host and deploy large language models (LLMs) for inference. No GPU, internet access or additional tools are needed to get started except for [Docker](https://docs.docker.com/desktop/install/linux-install/)!

AIKit uses [LocalAI](https://localai.io/) under-the-hood to run inference. LocalAI provides a drop-in replacement REST API that is OpenAI API compatible, so you can use any OpenAI API compatible client, such as [Kubectl AI](https://github.com/sozercan/kubectl-ai), [Chatbot-UI](https://github.com/sozercan/chatbot-ui) and many more, to send requests to open-source LLMs powered by AIKit!

👉 For full documentation, please see [AIKit website](https://sozercan.github.io/aikit/)!

## Features

- 🐳 No GPU, Internet access or additional tools needed except for [Docker](https://docs.docker.com/desktop/install/linux-install/)!
- 🤏 Minimal image size, resulting in less vulnerabilities and smaller attack surface with a custom [distroless](https://github.com/GoogleContainerTools/distroless)-based image
- 🚀 [Easy to use declarative configuration](https://sozercan.github.io/aikit/specs)
- ✨ OpenAI API compatible to use with any OpenAI API compatible client
- 📸 [Multi-modal model support](./docs/demo.md#vision-with-llava)
- 🖼️ Image generation support with Stable Diffusion
- 🦙 Support for GGUF ([`llama`](https://github.com/ggerganov/llama.cpp)), GPTQ ([`exllama`](https://github.com/turboderp/exllama) or [`exllama2`](https://github.com/turboderp/exllamav2)), EXL2 ([`exllama2`](https://github.com/turboderp/exllamav2)), and GGML ([`llama-ggml`](https://github.com/ggerganov/llama.cpp)) and [Mamba](https://github.com/state-spaces/mamba) models
- 🚢 [Kubernetes deployment ready](#kubernetes-deployment)
- 📦 Supports multiple models with a single image
- 🖥️ [Supports GPU-accelerated inferencing with NVIDIA GPUs](#nvidia)
- 🔐 [Signed images for `aikit` and pre-made models](https://sozercan.github.io/aikit/cosign)
- 🌈 Support for non-proprietary and self-hosted container registries to store model images

## Quick Start

You can get started with AIKit quickly on your local machine without a GPU!

```bash
docker run -d --rm -p 8080:8080 ghcr.io/sozercan/llama2:7b
```

```bash
curl http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d '{
    "model": "llama-2-7b-chat",
    "messages": [{"role": "user", "content": "explain kubernetes in a sentence"}]
  }'
```

Output should be similar to:

`{"created":1701236489,"object":"chat.completion","id":"dd1ff40b-31a7-4418-9e32-42151ab6875a","model":"llama-2-7b-chat","choices":[{"index":0,"finish_reason":"stop","message":{"role":"assistant","content":"\nKubernetes is a container orchestration system that automates the deployment, scaling, and management of containerized applications in a microservices architecture."}}],"usage":{"prompt_tokens":0,"completion_tokens":0,"total_tokens":0}}`

That's it! 🎉 API is OpenAI compatible so this is a drop-in replacement for any OpenAI API compatible client.

## Pre-made Models

AIKit comes with pre-made models that you can use out-of-the-box!

### CPU

| Model     | Optimization | Parameters | Command                                                         | License                                                                             |
| --------- | ------------ | ---------- | --------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| 🦙 Llama 2 | Chat         | 7B         | `docker run -d --rm -p 8080:8080 ghcr.io/sozercan/llama2:7b`    | [Llama 2](https://ai.meta.com/llama/license/)                                       |
| 🦙 Llama 2 | Chat         | 13B        | `docker run -d --rm -p 8080:8080 ghcr.io/sozercan/llama2:13b`   | [Llama 2](https://ai.meta.com/llama/license/)                                       |
| 🐬 Orca 2  |              | 13B        | `docker run -d --rm -p 8080:8080 ghcr.io/sozercan/orca2:13b`    | [Microsoft Research](https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE) |
| Ⓜ️ Mixtral | Instruct     | 8x7B       | `docker run -d --rm -p 8080:8080 ghcr.io/sozercan/mixtral:8x7b` | [Apache](https://choosealicense.com/licenses/apache-2.0/)                           |
| 🅿️ Phi 2   | Instruct     | 2.7B       | `docker run -d --rm -p 8080:8080 ghcr.io/sozercan/phi2:2.7b`    | [MIT](https://huggingface.co/microsoft/phi-2/resolve/main/LICENSE)                  |

### NVIDIA CUDA

| Model          | Optimization | Parameters | Command                                                                         | License                                                                             |
| -------------- | ------------ | ---------- | ------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| 🦙 Llama 2 Chat | Chat         | 7B         | `docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/llama2:7b-cuda`    | [Llama 2](https://ai.meta.com/llama/license/)                                       |
| 🦙 Llama 2 Chat | Chat         | 13B        | `docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/llama2:13b-cuda`   | [Llama 2](https://ai.meta.com/llama/license/)                                       |
| 🐬 Orca 2       |              | 13B        | `docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/orca2:13b-cuda`    | [Microsoft Research](https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE) |
| Ⓜ️ Mixtral      | Instruct     | 8x7B       | `docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/mixtral:8x7b-cuda` | [Apache](https://choosealicense.com/licenses/apache-2.0/)                           |
| 🅿️ Phi 2        | Instruct     | 2.7B       | `docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/phi2:2.7b-cuda`    | [MIT](https://huggingface.co/microsoft/phi-2/resolve/main/LICENSE)                  |

👉 For more information and how to create your own images, please see [AIKit website](https://sozercan.github.io/aikit/).

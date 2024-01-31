---
title: Introduction
slug: /
---

AIKit is a quick, easy, and local or cloud-agnostic way to get started to host and deploy large language models (LLMs) for inference. No GPU, internet access or additional tools are needed to get started except for [Docker](https://docs.docker.com/desktop/install/linux-install/)!

AIKit uses [LocalAI](https://localai.io/) under-the-hood to run inference. LocalAI provides a drop-in replacement REST API that is OpenAI API compatible, so you can use any OpenAI API compatible client, such as [Kubectl AI](https://github.com/sozercan/kubectl-ai), [Chatbot-UI](https://github.com/sozercan/chatbot-ui) and many more, to send requests to open-source LLMs powered by AIKit!

## Features

- 🐳 No GPU, Internet access or additional tools needed except for [Docker](https://docs.docker.com/desktop/install/linux-install/)!
- 🤏 Minimal image size, resulting in less vulnerabilities and smaller attack surface with a custom [distroless](https://github.com/GoogleContainerTools/distroless)-based image
- 🚀 [Easy to use declarative configuration](specs.md)
- ✨ OpenAI API compatible to use with any OpenAI API compatible client
- 📸 [Multi-modal model support](vision.md)
- 🖼️ Image generation support with Stable Diffusion
- 🦙 Support for GGUF ([`llama`](https://github.com/ggerganov/llama.cpp)), GPTQ ([`exllama`](https://github.com/turboderp/exllama) or [`exllama2`](https://github.com/turboderp/exllamav2)), EXL2 ([`exllama2`](https://github.com/turboderp/exllamav2)), and GGML ([`llama-ggml`](https://github.com/ggerganov/llama.cpp)) and [Mamba](https://github.com/state-spaces/mamba) models
- 🚢 [Kubernetes deployment ready](#kubernetes-deployment)
- 📦 Supports multiple models with a single image
- 🖥️ [Supports GPU-accelerated inferencing with NVIDIA GPUs](gpu.md)
- 🔐 [Signed images for `aikit` and pre-made models](cosign.md)
- 🌈 Support for non-proprietary and self-hosted container registries to store model images

To get started, please see [Quick Start](quick-start.md)!

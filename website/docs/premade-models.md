---
title: Pre-made Models
---

AIKit comes with pre-made models that you can use out-of-the-box!

If it doesn't include a specific model, you can always [create your own images](https://sozercan.github.io/aikit/premade-models/), and host in a container registry of your choice!

### CPU

| Model     | Optimization | Parameters | Command                                                         | Model Name              | License                                                                             |
| --------- | ------------ | ---------- | --------------------------------------------------------------- | ----------------------- | ----------------------------------------------------------------------------------- |
| 🦙 Llama 3 | Instruct     | 8B         | `docker run -d --rm -p 8080:8080 ghcr.io/sozercan/llama3:8b`    | `llama-3-8b-instruct`   | [Llama](https://ai.meta.com/llama/license/)                                         |
| 🦙 Llama 3 | Instruct     | 70B        | `docker run -d --rm -p 8080:8080 ghcr.io/sozercan/llama3:70b`   | `llama-3-70b-instruct`  | [Llama](https://ai.meta.com/llama/license/)                                         |
| 🦙 Llama 2 | Chat         | 7B         | `docker run -d --rm -p 8080:8080 ghcr.io/sozercan/llama2:7b`    | `llama-2-7b-chat`       | [Llama](https://ai.meta.com/llama/license/)                                         |
| 🦙 Llama 2 | Chat         | 13B        | `docker run -d --rm -p 8080:8080 ghcr.io/sozercan/llama2:13b`   | `llama-2-13b-chat`      | [Llama](https://ai.meta.com/llama/license/)                                         |
| Ⓜ️ Mixtral | Instruct     | 8x7B       | `docker run -d --rm -p 8080:8080 ghcr.io/sozercan/mixtral:8x7b` | `mixtral-8x7b-instruct` | [Apache](https://choosealicense.com/licenses/apache-2.0/)                           |
| 🅿️ Phi 3   | Instruct     | 3.8B       | `docker run -d --rm -p 8080:8080 ghcr.io/sozercan/phi3:3.8b`    | `phi-3-3.8b`            | [MIT](https://huggingface.co/microsoft/Phi-3-mini-4k-instruct/resolve/main/LICENSE) |

### NVIDIA CUDA

| Model     | Optimization | Parameters | Command                                                                         | Model Name              | License                                                                             |
| --------- | ------------ | ---------- | ------------------------------------------------------------------------------- | ----------------------- | ----------------------------------------------------------------------------------- |
| 🦙 Llama 3 | Instruct     | 8B         | `docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/llama3:8b-cuda`    | `llama-3-8b-instruct`   | [Llama](https://ai.meta.com/llama/license/)                                         |
| 🦙 Llama 3 | Instruct     | 70B        | `docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/llama3:70b-cuda`   | `llama-3-70b-instruct`  | [Llama](https://ai.meta.com/llama/license/)                                         |
| 🦙 Llama 2 | Chat         | 7B         | `docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/llama2:7b-cuda`    | `llama-2-7b-chat`       | [Llama](https://ai.meta.com/llama/license/)                                         |
| 🦙 Llama 2 | Chat         | 13B        | `docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/llama2:13b-cuda`   | `llama-2-13b-chat`      | [Llama](https://ai.meta.com/llama/license/)                                         |
| Ⓜ️ Mixtral | Instruct     | 8x7B       | `docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/mixtral:8x7b-cuda` | `mixtral-8x7b-instruct` | [Apache](https://choosealicense.com/licenses/apache-2.0/)                           |
| 🅿️ Phi 3   | Instruct     | 3.8B       | `docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/phi3:3.8b-cuda`    | `phi-3-3.8b`            | [MIT](https://huggingface.co/microsoft/Phi-3-mini-4k-instruct/resolve/main/LICENSE) |

:::note
Please see [models folder](https://github.com/sozercan/aikit/tree/main/models) for pre-made model definitions.

If not being offloaded to GPU VRAM, minimum of 8GB of RAM is required for 7B models, 16GB of RAM to run 13B models, and 32GB of RAM to run 8x7B models.

CPU models requires minimum of [AVX instruction set](https://en.wikipedia.org/wiki/Advanced_Vector_Extensions). You can check if your CPU supports AVX by running `grep avx /proc/cpuinfo`.

CUDA models includes CUDA v12. They are used with [NVIDIA GPU acceleration](#gpu-acceleration-support).
:::

## Deprecated Models

The following pre-made models are deprecated and no longer updated. Images will continue to be pullable, if needed.

If you need to use these specific models, you can always [create your own images](./create-images.md), and host in a container registry of your choice!

### CPU

| Model    | Optimization | Parameters | Command                                                      | License                                                                             |
| -------- | ------------ | ---------- | ------------------------------------------------------------ | ----------------------------------------------------------------------------------- |
| 🐬 Orca 2 |              | 13B        | `docker run -d --rm -p 8080:8080 ghcr.io/sozercan/orca2:13b` | [Microsoft Research](https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE) |
| 🅿️ Phi 2  | Instruct     | 2.7B       | `docker run -d --rm -p 8080:8080 ghcr.io/sozercan/phi2:2.7b` | [MIT](https://huggingface.co/microsoft/phi-2/resolve/main/LICENSE)                  |

### NVIDIA CUDA

| Model    | Optimization | Parameters | Command                                                                      | License                                                                             |
| -------- | ------------ | ---------- | ---------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| 🐬 Orca 2 |              | 13B        | `docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/orca2:13b-cuda` | [Microsoft Research](https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE) |
| 🅿️ Phi 2  | Instruct     | 2.7B       | `docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/phi2:2.7b-cuda` | [MIT](https://huggingface.co/microsoft/phi-2/resolve/main/LICENSE)                  |

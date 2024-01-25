---
title: llama.cpp (GGUF and GGML)
---

[Llama.cpp](https://github.com/ggerganov/llama.cpp) is a port of Facebook's LLaMA model in C/C++.

This is the default backend for `aikit`. No additional configuration is required.

This backend:
- provides support for GGUF (recommended) and GGML models
- supports both CPU (AVX, AVX2 or AVX512) and CUDA runtimes

## Example

:::warning
Please make sure to change syntax to `#syntax=ghcr.io/sozercan/aikit:latest` in the examples below.
:::

### CPU
https://github.com/sozercan/aikit/blob/main/test/aikitfile-llama.yaml

### GPU (CUDA)
https://github.com/sozercan/aikit/blob/main/test/aikitfile-llama-cuda.yaml

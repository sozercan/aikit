---
title: Exllama (GPTQ)
---

[Exllama](https://github.com/turboderp/exllama) is a standalone Python/C++/CUDA implementation of Llama for use with 4-bit GPTQ weights, designed to be fast and memory-efficient on modern GPUs.

This backend:
- provides support for GPTQ models
- requires CUDA runtime

:::note
This is an experimental backend and it may change in the future.
:::

## Example

:::warning
Please make sure to change syntax to `#syntax=ghcr.io/sozercan/aikit:latest` in the examples below.
:::

https://github.com/sozercan/aikit/blob/main/test/aikitfile-exllama.yaml

---
title: Diffusion
---

AIKit supports `diffusers` and `stablediffusion` backends.

## diffusers

`diffusers` backend uses the huggingface [`diffusers`](https://huggingface.co/docs/diffusers/en/index) library to generate images.


## stablediffusion

https://github.com/EdVince/Stable-Diffusion-NCNN

This backend:
- provides support for Stable Diffusion models
- does not support CUDA runtime yet

:::note
This is an experimental backend and it may change in the future.
:::

## Example

:::warning
Please make sure to change syntax to `#syntax=ghcr.io/sozercan/aikit:latest` in the examples below.
:::

https://github.com/sozercan/aikit/blob/main/test/aikitfile-stablediffusion.yaml

## Demo

https://www.youtube.com/watch?v=gh7b-rt70Ug

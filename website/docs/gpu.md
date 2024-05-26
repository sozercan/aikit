---
title: GPU Acceleration
---

:::note
At this time, only NVIDIA GPU acceleration is supported. Please open an issue if you'd like to see support for other GPU vendors.
:::

## NVIDIA

AIKit supports GPU accelerated inferencing with [NVIDIA Container Toolkit](https://github.com/NVIDIA/nvidia-container-toolkit). You must also have [NVIDIA Drivers](https://www.nvidia.com/en-us/drivers/unix/) installed on your host machine.

For Kubernetes, [NVIDIA GPU Operator](https://github.com/NVIDIA/gpu-operator) provides a streamlined way to install the NVIDIA drivers and container toolkit to configure your cluster to use GPUs.

To get started with GPU-accelerated inferencing, make sure to set the following in your `aikitfile` and build your model.

```yaml
runtime: cuda         # use NVIDIA CUDA runtime
```

For `llama` backend, set the following in your `config`:

```yaml
f16: true             # use float16 precision
gpu_layers: 35        # number of layers to offload to GPU
low_vram: true        # for devices with low VRAM
```

:::tip
Make sure to customize these values based on your model and GPU specs.
:::

:::note
For `exllama` and `exllama2` backends, GPU acceleration is enabled by default and cannot be disabled.
:::

After building the model, you can run it with [`--gpus all`](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/latest/docker-specialized.html#gpu-enumeration) flag to enable GPU support:

```bash
# for pre-made models, replace "my-model" with the image name
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

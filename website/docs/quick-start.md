---
title: Quick Start
---

You can get started with AIKit quickly on your local machine without a GPU!

```bash
docker run -d --rm -p 8080:8080 ghcr.io/sozercan/llama3:8b
```

```bash
curl http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d '{
    "model": "llama-3-8b-instruct",
    "messages": [{"role": "user", "content": "explain kubernetes in a sentence"}]
  }'
```

Output should be similar to:

```jsonc
{
  // ...
    "model": "llama-3-8b-instruct",
    "choices": [
        {
            "index": 0,
            "finish_reason": "stop",
            "message": {
                "role": "assistant",
                "content": "Kubernetes is an open-source container orchestration system that automates the deployment, scaling, and management of applications and services, allowing developers to focus on writing code rather than managing infrastructure."
            }
        }
    ],
    // ...
}
```

That's it! 🎉 API is OpenAI compatible so this is a drop-in replacement for any OpenAI API compatible client.

## What's next?

👉 If you are interested in other pre-made models (such as Mistral or Mixtral), please refer to [Pre-made models](./premade-models.md).

👉  If you are interested in learning more about how to create your own custom model images, please refer to [Creating Model Images](./create-images.md).

👉  If you are interested in fine tuning a model with domain-specific knowledge, please refer to [Fine Tuning](./fine-tune.md).

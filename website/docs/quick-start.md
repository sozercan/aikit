---
title: Quick Start
---

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
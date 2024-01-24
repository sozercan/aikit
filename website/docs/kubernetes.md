---
title: Kubernetes Deployment
---

It is easy to get started to deploy your models to Kubernetes!

Make sure you have a Kubernetes cluster running and `kubectl` is configured to talk to it, and your model images are accessible from the cluster.

:::tip
You can use [kind](https://kind.sigs.k8s.io/) to create a local Kubernetes cluster for testing purposes.
:::

```bash
# create a deployment
# for pre-made models, replace "my-model" with the image name
kubectl create deployment my-llm-deployment --image=my-model

# expose it as a service
kubectl expose deployment my-llm-deployment --port=8080 --target-port=8080 --name=my-llm-service

# easy to scale up and down as needed
kubectl scale deployment my-llm-deployment --replicas=3

# port-forward for testing locally
kubectl port-forward service/my-llm-service 8080:8080

# send requests to your model
curl http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d '{
     "model": "llama-2-7b-chat",
     "messages": [{"role": "user", "content": "explain kubernetes in a sentence"}]
   }'
{"created":1701236489,"object":"chat.completion","id":"dd1ff40b-31a7-4418-9e32-42151ab6875a","model":"llama-2-7b-chat","choices":[{"index":0,"finish_reason":"stop","message":{"role":"assistant","content":"\nKubernetes is a container orchestration system that automates the deployment, scaling, and management of containerized applications in a microservices architecture."}}],"usage":{"prompt_tokens":0,"completion_tokens":0,"total_tokens":0}}
```

:::tip
For an example Kubernetes deployment and service YAML, see [kubernetes folder](./kubernetes/). Please note that these are examples, you may need to customize them (such as properly configured resource requests and limits) based on your needs.
:::

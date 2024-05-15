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

## Helm Chart

For more advanced deployments, you can use the Helm chart provided in the `charts` directory.

Install the chart using the following command:

```bash
helm install aikit ./charts/aikit
```

Output will be similar to:

```bash
NAME: aikit
LAST DEPLOYED: Wed May 15 05:32:39 2024
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
Get the application URL by running these commands:
  echo "Visit http://127.0.0.1:8080 to access aikit WebUI."
  kubectl --namespace default port-forward service/aikit-webui 8080:80
```

As mentioned in the notes, you can then port-forward and hhen navigate to the URL provided to access the WebUI.

### Values
| Key                | Type   | Default          | Description       |
| ------------------ | ------ | ---------------- | ----------------- |
| `model.name`       | string | `"model"`        | Model name        |
| `image.repository` | string | `"aikit"`        | Image repository  |
| `image.tag`        | string | `"latest"`       | Image tag         |
| `image.pullPolicy` | string | `"IfNotPresent"` | Image pull policy |
| `service.type`     | string | `"ClusterIP"`    | Service type      |
```

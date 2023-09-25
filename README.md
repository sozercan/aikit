# AIKit ✨

AIKit is a quick and easy way to get started to host and deploy large language models (LLMs) for inference. No GPU or additional tools needed except for Docker! 🐳

AIKit uses [LocalAI](https://localai.io/) under the hood to run inference. LocalAI is OpenAI API compatible, so you can use any OpenAI API compatible client, such as [Kubectl AI](https://github.com/sozercan/kubectl-ai), to send requests to open-source LLMs powered by AIKit!

## Demo

[![asciicast]()]()

## Getting Started

Create an `aikitfile.yaml` with the following structure:

```yaml
#syntax=sozercan/aikit:latest
apiVersion: v1alpha1
models:
  - name: mistral-7b
    source: https://huggingface.co/TheBloke/Mistral-7B-OpenOrca-GGUF/resolve/main/mistral-7b-openorca.Q6_K.gguf
```

> For full API reference, see [API](docs/api.md).

> You can find more models at [model gallery](https://github.com/go-skynet/model-gallery).

Then run the following command:

```bash
docker build -t my-model .
```

> If you name your file something else, you can pass it to `docker build` command with `--file` flag.

This will build a docker image with your models. You can then run it with:

```bash
docker run -p 8080:8080 my-model
```

You can then send requests to `localhost:8080` to run inference from your models.
```bash
curl http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d '{
     "model": "mistral-7b-openorca.Q6_K.gguf",
     "messages": [{"role": "user", "content": "explain kubernetes in less than 100 words"}],
     "temperature": 0.5
   }'
{"created":1700542409,"object":"chat.completion","id":"64f8b8e6-328d-4d24-8fc2-4b7b45fbdcc0","model":"mistral-7b-openorca.Q6_K.gguf","choices":[{"index":0,"finish_reason":"stop","message":{"role":"assistant","content":"\n\nKubernetes is an open-source platform for automating deployment, scaling, and management of containerized applications. It provides a way to manage and orchestrate containers across multiple hosts, ensuring that applications are always available and running efficiently.\n\nKubernetes is like a traffic cop for containers, making sure they are placed where they need to be, have the resources they need, and are working together properly. It helps developers and operations teams manage complex containerized environments with ease."}}],"usage":{"prompt_tokens":0,"completion_tokens":0,"total_tokens":0}}
```

## Kubernetes Deployment

It is easy to get started to deploy your models to Kubernetes!

Make sure you have a Kubernetes cluster running and `kubectl` is configured to talk to it, and your model images are accessible from the cluster. You can use [kind](https://kind.sigs.k8s.io/) to create a local Kubernetes cluster for testing purposes.

```bash
# create a test cluster using kind
kind create cluster

# load your local image to the cluster
kind load docker-image my-model

# create a deployment
kubectl create deployment my-llm-deployment --image=my-model --image-pull-policy=IfNotPresent

# expose it as a service
kubectl expose deployment my-llm-deployment --port=8080 --target-port=8008 --name=my-llm-service

# easy to scale up and down
kubectl scale deployment my-llm-deployment --replicas=3

# port-forward for testing locally
kubectl port-forward service/my-llm-service 8080:8080

# send requests to your model
curl http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d '{
     "model": "mistral-7b-openorca.Q6_K.gguf",
     "messages": [{"role": "user", "content": "explain kubernetes in less than 100 words"}],
     "temperature": 0.5
   }'
{"created":1700542409,"object":"chat.completion","id":"64f8b8e6-328d-4d24-8fc2-4b7b45fbdcc0","model":"mistral-7b-openorca.Q6_K.gguf","choices":[{"index":0,"finish_reason":"stop","message":{"role":"assistant","content":"\n\nKubernetes is an open-source platform for automating deployment, scaling, and management of containerized applications. It provides a way to manage and orchestrate containers across multiple hosts, ensuring that applications are always available and running efficiently.\n\nKubernetes is like a traffic cop for containers, making sure they are placed where they need to be, have the resources they need, and are working together properly. It helps developers and operations teams manage complex containerized environments with ease."}}],"usage":{"prompt_tokens":0,"completion_tokens":0,"total_tokens":0}}
```

> For an example Kubernetes deployment YAML, see [deployment.yaml](kubernetes/deployment.yaml).

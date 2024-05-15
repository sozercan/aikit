# Helm chart

This directory contains the Helm chart for deploying AIKit on Kubernetes.

## Usage


## Values
| Key                | Type   | Default          | Description       |
| ------------------ | ------ | ---------------- | ----------------- |
| `model.name`       | string | `"model"`        | Model name        |
| `image.repository` | string | `"aikit"`        | Image repository  |
| `image.tag`        | string | `"latest"`       | Image tag         |
| `image.pullPolicy` | string | `"IfNotPresent"` | Image pull policy |
| `service.type`     | string | `"ClusterIP"`    | Service type      |

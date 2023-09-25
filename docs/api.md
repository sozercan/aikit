# API

## v1alpha1

```go
type Config struct {
    // API version (v1alpha1)
	ApiVersion string  `yaml:"apiVersion"`
    // Debug mode for the LocalAI server (default: false)
    Debug      bool    `yaml:"debug"`
    // Model definition list
	Models     []Model `yaml:"models"`
}

type Model struct {
    // name of the model
	Name   string `yaml:"name"`
    // source of the model from a URL
	Source string `yaml:"source"`
    // Config file
	Config string `yaml:"config"`
}
```

Example:

```yaml
apiVersion: v1alpha1
models:
  - name: mistral-7b
    source: https://huggingface.co/TheBloke/Mistral-7B-OpenOrca-GGUF/resolve/main/mistral-7b-openorca.Q6_K.gguf
```


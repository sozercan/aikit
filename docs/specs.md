# API Specifications

## v1alpha1

```go
type Config struct {
	APIVersion string  `yaml:"apiVersion"`
	Debug      bool    `yaml:"debug,omitempty"`
	Runtime    string  `yaml:"runtime,omitempty"`
	Models     []Model `yaml:"models"`
	Config     string  `yaml:"config,omitempty"`
}

type Model struct {
	Name            string           `yaml:"name"`
	Source          string           `yaml:"source"`
	SHA256          string           `yaml:"sha256,omitempty"`
	PromptTemplates []PromptTemplate `yaml:"promptTemplates,omitempty"`
}

type PromptTemplate struct {
	Name     string `yaml:"name,omitempty"`
	Template string `yaml:"template,omitempty"`
}
```

Example:

```yaml
#syntax=docker.io/sozercan/aikit:latest
apiVersion: v1alpha1
debug: true
runtime: gpu-nvidia
models:
  - name: llama-2-7b-chat
    source: https://huggingface.co/TheBloke/Llama-2-7B-Chat-GGUF/resolve/main/llama-2-7b-chat.Q4_K_M.gguf
    sha256: "08a5566d61d7cb6b420c3e4387a39e0078e1f2fe5f055f3a03887385304d4bfa"
    promptTemplates:
      - name: "llama-2-7b-chat"
        template: |
          {{if eq .RoleName \"assistant\"}}{{.Content}}{{else}}
          [INST]
          {{if .SystemPrompt}}{{.SystemPrompt}}{{else if eq .RoleName \"system\"}}<<SYS>>{{.Content}}<</SYS>>

          {{else if .Content}}{{.Content}}{{end}}
          [/INST]
          {{end}}
config: |
  - name: \"llama-2-7b-chat\"
    backend: \"llama\"
    parameters:
      top_k: 80
      temperature: 0.2
      top_p: 0.7
      model: \"llama-2-7b-chat.Q4_K_M.gguf\"
    context_size: 4096
    roles:
      function: 'Function Result:'
      assistant_function_call: 'Function Call:'
      assistant: 'Assistant:'
      user: 'User:'
      system: 'System:'
    template:
      chat_message: \"llama-2-7b-chat\"
    system_prompt: \"You are a helpful assistant, below is a conversation, please respond with the next message and do not ask follow-up questions\"
```

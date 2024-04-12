//go:generate go run ../../../cmd/gen-jsonschema ../../../schema/inferencespec.schema.json
package config

type InferenceConfig struct {
	APIVersion string   `yaml:"apiVersion" json:"apiVersion" jsonschema:"required"`
	Debug      bool     `yaml:"debug,omitempty" json:"debug,omitempty"`
	Runtime    string   `yaml:"runtime,omitempty" json:"runtime,omitempty"`
	Backends   []string `yaml:"backends,omitempty" json:"backends,omitempty"`
	Models     []Model  `yaml:"models" json:"models" jsonschema:"required"`
	Config     string   `yaml:"config,omitempty" json:"config,omitempty"`
}

type Model struct {
	Name            string           `yaml:"name" json:"name" jsonschema:"required"`
	Source          string           `yaml:"source" json:"source" jsonschema:"required"`
	SHA256          string           `yaml:"sha256,omitempty" json:"sha256,omitempty"`
	PromptTemplates []PromptTemplate `yaml:"promptTemplates,omitempty" json:"promptTemplates,omitempty"`
}

type PromptTemplate struct {
	Name     string `yaml:"name,omitempty" json:"name,omitempty"`
	Template string `yaml:"template,omitempty" json:"template,omitempty"`
}

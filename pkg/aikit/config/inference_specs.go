//go:generate go run ../../../cmd/gen-jsonschema ../../../schema/inferencespec.schema.json
package config

type InferenceConfig struct {
	APIVersion string   `yaml:"apiVersion" json:"apiVersion" jsonschema:"required"`
	Debug      bool     `yaml:"debug" json:"debug"`
	Runtime    string   `yaml:"runtime" json:"runtime"`
	Backends   []string `yaml:"backends" json:"backends"`
	Models     []Model  `yaml:"models" json:"models" jsonschema:"required"`
	Config     string   `yaml:"config" json:"config"`
}

type Model struct {
	Name            string           `yaml:"name" json:"name"`
	Source          string           `yaml:"source" json:"source"`
	SHA256          string           `yaml:"sha256" json:"sha256"`
	PromptTemplates []PromptTemplate `yaml:"promptTemplates" json:"promptTemplates"`
}

type PromptTemplate struct {
	Name     string `yaml:"name" json:"name"`
	Template string `yaml:"template" json:"template"`
}

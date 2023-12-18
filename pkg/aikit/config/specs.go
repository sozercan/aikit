package config

import (
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

func NewFromBytes(b []byte) (*Config, error) {
	c := &Config{}
	if err := yaml.Unmarshal(b, c); err != nil {
		return nil, errors.Wrap(err, "unmarshal config")
	}
	return c, nil
}

type Config struct {
	APIVersion string   `yaml:"apiVersion"`
	Debug      bool     `yaml:"debug,omitempty"`
	Runtime    string   `yaml:"runtime,omitempty"`
	Backends   []string `yaml:"backends,omitempty"`
	Models     []Model  `yaml:"models"`
	Config     string   `yaml:"config,omitempty"`
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

package config

import (
	"io"
	"os"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// NewFromFilename returns a new config from a filename.
func NewFromFilename(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "opening file")
	}
	defer f.Close()
	contents, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "reading config file")
	}
	return NewFromBytes(contents)
}

func NewFromBytes(b []byte) (*Config, error) {
	c := &Config{}
	if err := yaml.Unmarshal(b, c); err != nil {
		return nil, errors.Wrap(err, "unmarshal config")
	}
	return c, nil
}

type Config struct {
	APIVersion string  `yaml:"apiVersion"`
	Debug      bool    `yaml:"debug"`
	Models     []Model `yaml:"models"`
}

type Model struct {
	Name   string `yaml:"name"`
	Source string `yaml:"source"`
	Config string `yaml:"config"`
}

package config

import (
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

func NewFromBytes(b []byte) (*InferenceConfig, *FineTuneConfig, error) {
	inferenceConfig := &InferenceConfig{}
	fineTuneConfig := &FineTuneConfig{}
	var err error
	err = yaml.Unmarshal(b, inferenceConfig)
	if err == nil {
		return inferenceConfig, nil, nil
	}

	err = yaml.Unmarshal(b, fineTuneConfig)
	if err == nil {
		return nil, fineTuneConfig, nil
	}

	return nil, nil, errors.Wrap(err, "unmarshal config")
}

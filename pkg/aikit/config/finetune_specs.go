//go:generate go run ../../../cmd/gen-jsonschema ../../../schema/finetunespec.schema.json
package config

type FineTuneConfig struct {
	APIVersion string             `yaml:"apiVersion" json:"apiVersion" jsonschema:"required"`
	Target     string             `yaml:"target" json:"target"`
	BaseModel  string             `yaml:"baseModel" json:"baseModel" jsonschema:"required"`
	Datasets   []Dataset          `yaml:"datasets" json:"datasets" jsonschema:"required"`
	Config     FineTuneConfigSpec `yaml:"config" json:"config"`
	Output     FineTuneOutputSpec `yaml:"output" json:"output"`
}

type FineTuneConfigSpec struct {
	Unsloth FineTuneConfigUnslothSpec `yaml:"unsloth" json:"unsloth"`
}

type Dataset struct {
	Source string `yaml:"source" json:"source"`
	Type   string `yaml:"type" json:"type"`
}

type FineTuneConfigUnslothSpec struct {
	Packing                   bool    `yaml:"packing" json:"packing"`
	MaxSeqLength              int     `yaml:"maxSeqLength" json:"maxSeqLength"`
	LoadIn4bit                bool    `yaml:"loadIn4bit" json:"loadIn4bit"`
	BatchSize                 int     `yaml:"batchSize" json:"batchSize"`
	GradientAccumulationSteps int     `yaml:"gradientAccumulationSteps" json:"gradientAccumulationSteps"`
	WarmupSteps               int     `yaml:"warmupSteps" json:"warmupSteps"`
	MaxSteps                  int     `yaml:"maxSteps" json:"maxSteps"`
	LearningRate              float64 `yaml:"learningRate" json:"learningRate"`
	LoggingSteps              int     `yaml:"loggingSteps" json:"loggingSteps"`
	Optimizer                 string  `yaml:"optimizer" json:"optimizer"`
	WeightDecay               float64 `yaml:"weightDecay" json:"weightDecay"`
	LrSchedulerType           string  `yaml:"lrSchedulerType" json:"lrSchedulerType"`
	Seed                      int     `yaml:"seed" json:"seed"`
}

type FineTuneOutputSpec struct {
	Quantize string `yaml:"quantize" json:"quantize"`
	Name     string `yaml:"name" json:"name"`
}

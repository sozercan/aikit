package config

type FineTuneConfig struct {
	APIVersion string             `yaml:"apiVersion"`
	Target     string             `yaml:"target"`
	BaseModel  string             `yaml:"baseModel"`
	Datasets   []Dataset          `yaml:"datasets"`
	Config     FineTuneConfigSpec `yaml:"config"`
	Output     FineTuneOutputSpec `yaml:"output"`
}

type FineTuneConfigSpec struct {
	Unsloth FineTuneConfigUnslothSpec `yaml:"unsloth"`
}

type Dataset struct {
	Source string `yaml:"source"`
	Type   string `yaml:"type"`
}

type FineTuneConfigUnslothSpec struct {
	Packing                   bool    `yaml:"packing"`
	MaxSeqLength              int     `yaml:"maxSeqLength"`
	LoadIn4bit                bool    `yaml:"loadIn4bit"`
	BatchSize                 int     `yaml:"batchSize"`
	GradientAccumulationSteps int     `yaml:"gradientAccumulationSteps"`
	WarmupSteps               int     `yaml:"warmupSteps"`
	MaxSteps                  int     `yaml:"maxSteps"`
	LearningRate              float64 `yaml:"learningRate"`
	LoggingSteps              int     `yaml:"loggingSteps"`
	Optimizer                 string  `yaml:"optimizer"`
	WeightDecay               float64 `yaml:"weightDecay"`
	LrSchedulerType           string  `yaml:"lrSchedulerType"`
	Seed                      int     `yaml:"seed"`
}

type FineTuneOutputSpec struct {
	Quantize string `yaml:"quantize"`
	Name     string `yaml:"name"`
}

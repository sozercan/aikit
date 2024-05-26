package build

import (
	"reflect"
	"testing"

	"github.com/sozercan/aikit/pkg/aikit/config"
)

func Test_validateConfig(t *testing.T) {
	type args struct {
		c *config.InferenceConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "no config",
			args:    args{c: &config.InferenceConfig{}},
			wantErr: true,
		},
		{
			name: "unsupported api version",
			args: args{c: &config.InferenceConfig{
				APIVersion: "v10",
			}},
			wantErr: true,
		},
		{
			name: "invalid runtime",
			args: args{c: &config.InferenceConfig{
				APIVersion: "v1",
				Runtime:    "foo",
			}},
			wantErr: true,
		},
		{
			name: "no models",
			args: args{c: &config.InferenceConfig{
				APIVersion: "v1alpha1",
			}},
			wantErr: true,
		},
		{
			name: "valid backend",
			args: args{c: &config.InferenceConfig{
				APIVersion: "v1alpha1",
				Runtime:    "cuda",
				Backends:   []string{"exllama"},
				Models: []config.Model{
					{
						Name:   "test",
						Source: "foo",
					},
				},
			}},
			wantErr: false,
		},
		{
			name: "invalid backend",
			args: args{c: &config.InferenceConfig{
				APIVersion: "v1alpha1",
				Backends:   []string{"foo"},
				Models: []config.Model{
					{
						Name:   "test",
						Source: "foo",
					},
				},
			}},
			wantErr: true,
		},
		{
			name: "valid backend but no cuda runtime",
			args: args{c: &config.InferenceConfig{
				APIVersion: "v1alpha1",
				Backends:   []string{"exllama"},
				Models: []config.Model{
					{
						Name:   "test",
						Source: "foo",
					},
				},
			}},
			wantErr: true,
		},
		{
			name: "invalid backend combination 1",
			args: args{c: &config.InferenceConfig{
				APIVersion: "v1alpha1",
				Runtime:    "cuda",
				Backends:   []string{"exllama", "exllama2"},
				Models: []config.Model{
					{
						Name:   "test",
						Source: "foo",
					},
				},
			}},
			wantErr: true,
		},
		{
			name: "invalid backend combination 2",
			args: args{c: &config.InferenceConfig{
				APIVersion: "v1alpha1",
				Runtime:    "cuda",
				Backends:   []string{"exllama", "stablediffusion"},
				Models: []config.Model{
					{
						Name:   "test",
						Source: "foo",
					},
				},
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateInferenceConfig(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateFineTuneConfig(t *testing.T) {
	type args struct {
		c *config.FineTuneConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "no config",
			args:    args{c: &config.FineTuneConfig{}},
			wantErr: true,
		},
		{
			name: "unsupported api version",
			args: args{c: &config.FineTuneConfig{
				APIVersion: "v10",
			}},
			wantErr: true,
		},
		{
			name: "invalid target",
			args: args{c: &config.FineTuneConfig{
				APIVersion: "v1alpha1",
				Target:     "foo",
			}},
			wantErr: true,
		},
		{
			name: "no datasets",
			args: args{c: &config.FineTuneConfig{
				APIVersion: "v1alpha1",
				Target:     "unsloth",
			}},
			wantErr: true,
		},
		{
			name: "invalid dataset type",
			args: args{c: &config.FineTuneConfig{
				APIVersion: "v1alpha1",
				Target:     "unsloth",
				Datasets: []config.Dataset{
					{
						Source: "foo",
						Type:   "bar",
					},
				},
			}},
			wantErr: true,
		},
		{
			name: "valid dataset type",
			args: args{c: &config.FineTuneConfig{
				APIVersion: "v1alpha1",
				Target:     "unsloth",
				Datasets: []config.Dataset{
					{
						Source: "foo",
						Type:   "alpaca",
					},
				},
			}},
			wantErr: false,
		},
		{
			name: "multiple datasets",
			args: args{c: &config.FineTuneConfig{
				APIVersion: "v1alpha1",
				Target:     "unsloth",
				Datasets: []config.Dataset{
					{
						Source: "foo",
						Type:   "alpaca",
					},
					{
						Source: "bar",
						Type:   "alpaca",
					},
				},
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateFinetuneConfig(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateFineTuneConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defaultsUnslothConfig(t *testing.T) {
	type args struct {
		c *config.FineTuneConfig
	}
	tests := []struct {
		name string
		args args
		want *config.FineTuneConfig
	}{
		{
			name: "no config",
			args: args{c: &config.FineTuneConfig{}},
			want: &config.FineTuneConfig{
				Config: config.FineTuneConfigSpec{
					Unsloth: config.FineTuneConfigUnslothSpec{
						Packing:                   false,
						MaxSeqLength:              2048,
						LoadIn4bit:                false,
						BatchSize:                 2,
						GradientAccumulationSteps: 4,
						WarmupSteps:               10,
						MaxSteps:                  60,
						LearningRate:              0.0002,
						LoggingSteps:              1,
						Optimizer:                 "adamw_8bit",
						WeightDecay:               0.01,
						LrSchedulerType:           "linear",
						Seed:                      42,
					},
				},
			},
		},
		{
			name: "with config",
			args: args{c: &config.FineTuneConfig{
				Config: config.FineTuneConfigSpec{
					Unsloth: config.FineTuneConfigUnslothSpec{
						Packing:                   true,
						MaxSeqLength:              1024,
						LoadIn4bit:                true,
						BatchSize:                 4,
						GradientAccumulationSteps: 8,
						WarmupSteps:               20,
						MaxSteps:                  120,
						LearningRate:              0.0004,
						LoggingSteps:              2,
						Optimizer:                 "adamw_16bit",
						WeightDecay:               0.02,
						LrSchedulerType:           "cosine",
						Seed:                      24,
					},
				},
			}},
			want: &config.FineTuneConfig{
				Config: config.FineTuneConfigSpec{
					Unsloth: config.FineTuneConfigUnslothSpec{
						Packing:                   true,
						MaxSeqLength:              1024,
						LoadIn4bit:                true,
						BatchSize:                 4,
						GradientAccumulationSteps: 8,
						WarmupSteps:               20,
						MaxSteps:                  120,
						LearningRate:              0.0004,
						LoggingSteps:              2,
						Optimizer:                 "adamw_16bit",
						WeightDecay:               0.02,
						LrSchedulerType:           "cosine",
						Seed:                      24,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defaultsUnslothConfig(tt.args.c)
			if !reflect.DeepEqual(tt.args.c, tt.want) {
				t.Errorf("defaultsUnslothConfig() = %v, want %v", tt.args.c, tt.want)
			}
		})
	}
}

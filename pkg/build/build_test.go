package build

import (
	"testing"

	"github.com/sozercan/aikit/pkg/aikit/config"
)

func Test_validateConfig(t *testing.T) {
	type args struct {
		c *config.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "no config",
			args:    args{c: &config.Config{}},
			wantErr: true,
		},
		{
			name: "unsupported api version",
			args: args{c: &config.Config{
				APIVersion: "v10",
			}},
			wantErr: true,
		},
		{
			name: "invalid runtime",
			args: args{c: &config.Config{
				APIVersion: "v1",
				Runtime:    "foo",
			}},
			wantErr: true,
		},
		{
			name: "no models",
			args: args{c: &config.Config{
				APIVersion: "v1alpha1",
			}},
			wantErr: true,
		},
		{
			name: "valid backend",
			args: args{c: &config.Config{
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
			args: args{c: &config.Config{
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
			args: args{c: &config.Config{
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
			args: args{c: &config.Config{
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
			args: args{c: &config.Config{
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
			if err := validateConfig(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

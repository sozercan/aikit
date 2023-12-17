package aikit2llb

import (
	"github.com/moby/buildkit/util/system"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sozercan/aikit/pkg/aikit/config"
	"github.com/sozercan/aikit/pkg/utils"
)

func NewImageConfig(c *config.Config) *specs.Image {
	img := emptyImage(c)
	var debug, config string
	if c.Debug {
		debug = "--debug"
	}
	if c.Config != "" {
		config = "--config-file=/config.yaml"
	}
	img.Config.Entrypoint = []string{"local-ai", debug, config}
	return img
}

func emptyImage(c *config.Config) *specs.Image {
	img := &specs.Image{
		Platform: specs.Platform{
			Architecture: "amd64",
			OS:           "linux",
		},
	}
	img.RootFS.Type = "layers"
	img.Config.WorkingDir = "/"

	img.Config.Env = []string{
		"PATH=" + system.DefaultPathEnv("linux"),
	}

	cudaEnv := []string{
		"PATH=" + system.DefaultPathEnv("linux") + ":/usr/local/cuda/bin",
		"NVIDIA_REQUIRE_CUDA=cuda>=12.0",
		"NVIDIA_DRIVER_CAPABILITIES=compute,utility",
		"NVIDIA_VISIBLE_DEVICES=all",
		"LD_LIBRARY_PATH=/usr/local/cuda/lib64",
	}
	if c.Runtime == utils.RuntimeNVIDIA {
		img.Config.Env = append(img.Config.Env, cudaEnv...)
	}

	for b := range c.Backends {
		if c.Backends[b] == utils.BackendExllama {
			exllamaEnv := []string{
				"EXTERNAL_GRPC_BACKENDS=exllama:/tmp/localai/backend/python/exllama/exllama.py",
				"PYTHONPATH=/usr/local/cuda/lib64",
			}
			img.Config.Env = append(img.Config.Env, exllamaEnv...)
		}
	}

	return img
}

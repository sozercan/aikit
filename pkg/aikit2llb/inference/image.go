package inference

import (
	"github.com/moby/buildkit/util/system"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sozercan/aikit/pkg/aikit/config"
	"github.com/sozercan/aikit/pkg/utils"
)

func NewImageConfig(c *config.InferenceConfig) *specs.Image {
	img := emptyImage(c)
	cmd := []string{}
	if c.Debug {
		cmd = append(cmd, "--debug")
	}
	if c.Config != "" {
		cmd = append(cmd, "--config-file=/config.yaml")
	}

	img.Config.Entrypoint = []string{"local-ai"}
	img.Config.Cmd = cmd
	return img
}

func emptyImage(c *config.InferenceConfig) *specs.Image {
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
		switch c.Backends[b] {
		case utils.BackendExllama, utils.BackendExllamaV2:
			exllamaEnv := []string{
				"EXTERNAL_GRPC_BACKENDS=exllama:/tmp/localai/backend/python/exllama/exllama.py,exllama2:/tmp/localai/backend/python/exllama2/exllama2_backend.py",
				"CUDA_HOME=/usr/local/cuda",
			}
			img.Config.Env = append(img.Config.Env, exllamaEnv...)
		case utils.BackendMamba:
			mambaEnv := []string{
				"EXTERNAL_GRPC_BACKENDS=mamba:/tmp/localai/backend/python/mamba/backend_mamba.py",
				"CUDA_HOME=/usr/local/cuda",
			}
			img.Config.Env = append(img.Config.Env, mambaEnv...)
		}
	}

	return img
}

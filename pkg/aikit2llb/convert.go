package aikit2llb

import (
	"fmt"
	"net/url"
	"path"

	"github.com/moby/buildkit/client/llb"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sozercan/aikit/pkg/aikit/config"
)

const (
	debianSlim     = "docker.io/library/debian:12-slim"
	distrolessBase = "gcr.io/distroless/cc-debian12:latest"
	localAIVersion = "v1.40.0"
	retryCount     = 5
	cudaVersion    = "12-3"

	runtimeNVIDIA    = "gpu-nvidia"
	runtimeCPUAVX    = "cpu-avx"
	runtimeCPUAVX2   = "cpu-avx2"
	runtimeCPUAVX512 = "cpu-avx512"
)

func Aikit2LLB(c *config.Config) (llb.State, *specs.Image) {
	var merge llb.State
	s := llb.Image(debianSlim)
	s = curl(s)
	// if c.Runtime == runtimeNVIDIA {
	s, merge = installCuda(s)
	// } else {
	// 	merge = llb.Image(distrolessBase)
	// }
	s, merge = copyModels(s, merge, c)
	s = addLocalAI(c, s, merge)
	imageCfg := NewImageConfig(c)
	return s, imageCfg
}

func copyModels(s llb.State, merge llb.State, c *config.Config) (llb.State, llb.State) {
	initState := s

	// create config file if defined
	if c.Config != "" {
		s = s.Run(shf("echo \"%s\" >> /config.yaml", c.Config)).Root()
	}

	for _, model := range c.Models {
		s = s.Run(llb.Shlexf("curl --retry %d --create-dirs -sSLO --output-dir /models %s", retryCount, model.Source)).Root()
		// verify sha256 checksum if defined
		if model.SHA256 != "" {
			path := fmt.Sprintf("/models/%s", fileNameFromURL(model.Source))
			s = s.Run(shf("echo \"%s  %s\" | sha256sum -c -", model.SHA256, path)).Root()
		}
		// create prompt templates if defined
		for _, pt := range model.PromptTemplates {
			if pt.Name != "" && pt.Template != "" {
				s = s.Run(shf("echo \"%s\" >> /models/%s.tmpl", pt.Template, pt.Name)).Root()
			}
		}
	}
	diff := llb.Diff(initState, s)
	merge = llb.Merge([]llb.State{merge, diff})
	return s, merge
}

func fileNameFromURL(urlString string) string {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	return path.Base(parsedURL.Path)
}

func curl(s llb.State) llb.State {
	i := s.Run(llb.Shlex("apt-get update"), llb.IgnoreCache).Root()
	return i.Run(llb.Shlex("apt-get install curl -y")).Root()
}

func installCuda(s llb.State) (llb.State, llb.State) {
	initState := s

	s = s.Run(shf("curl -O https://developer.download.nvidia.com/compute/cuda/repos/debian12/x86_64/cuda-keyring_1.1-1_all.deb && dpkg -i cuda-keyring_1.1-1_all.deb && rm cuda-keyring_1.1-1_all.deb")).Root()
	s = s.Run(llb.Shlex("apt-get update"), llb.IgnoreCache).Root()
	s = s.Run(shf("apt-get install -y libcublas-%[1]s cuda-cudart-%[1]s && apt-get clean", cudaVersion)).Root()

	diff := llb.Diff(initState, s)
	merge := llb.Merge([]llb.State{llb.Image(distrolessBase), diff})
	return s, merge
}

func addLocalAI(c *config.Config, s llb.State, merge llb.State) llb.State {
	initState := s
	var localAIURL string
	switch c.Runtime {
	case runtimeNVIDIA:
		localAIURL = fmt.Sprintf("https://sertacstorage.blob.core.windows.net/localai/%s/local-ai", localAIVersion)
	case runtimeCPUAVX2:
		localAIURL = fmt.Sprintf("https://github.com/mudler/LocalAI/releases/download/%s/local-ai-avx2-Linux-x86_64", localAIVersion)
	case runtimeCPUAVX512:
		localAIURL = fmt.Sprintf("https://github.com/mudler/LocalAI/releases/download/%s/local-ai-avx512-Linux-x86_64", localAIVersion)
	case runtimeCPUAVX, "":
		localAIURL = fmt.Sprintf("https://github.com/mudler/LocalAI/releases/download/%s/local-ai-avx-Linux-x86_64", localAIVersion)
	}

	s = s.Run(llb.Shlexf("curl -Lo /usr/bin/local-ai %s", localAIURL)).Root()
	s = s.Run(llb.Shlex("chmod +x /usr/bin/local-ai")).Root()
	diff := llb.Diff(initState, s)
	return llb.Merge([]llb.State{merge, diff})
}

func shf(cmd string, v ...interface{}) llb.RunOption {
	return llb.Args([]string{"/bin/sh", "-c", fmt.Sprintf(cmd, v...)})
}

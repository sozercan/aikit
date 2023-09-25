package aikit2llb

import (
	"github.com/moby/buildkit/client/llb"
	"github.com/sozercan/aikit/pkg/aikit/config"
)

const (
	debianSlim     = "docker.io/library/debian:12-slim"
	distrolessBase = "gcr.io/distroless/cc-debian12:latest"
)

func Aikit2LLB(c *config.Config) (llb.State, *Image) {
	var merge llb.State
	s := llb.Image(debianSlim)
	s = curl(s)
	s, merge = copyModels(s, c)
	s = addLocalAI(s, merge)
	imageCfg := NewImageConfig(c)
	return s, imageCfg
}

func copyModels(s llb.State, c *config.Config) (llb.State, llb.State) {
	initState := s
	for _, model := range c.Models {
		s = s.Run(llb.Shlexf("curl --create-dirs -LO --output-dir /models %s", model.Source)).Root()
	}
	diff := llb.Diff(initState, s)
	merge := llb.Merge([]llb.State{llb.Image(distrolessBase), diff})
	return s, merge
}

func curl(s llb.State) llb.State {
	i := s.Run(llb.Shlex("apt update"), llb.IgnoreCache).Root()
	return i.Run(llb.Shlex("apt install curl -y")).Root()
}

func addLocalAI(s llb.State, merge llb.State) llb.State {
	initState := s

	localAIURL := "https://github.com/mudler/LocalAI/releases/download/v1.40.0/local-ai-avx-Linux-x86_64"
	c := s.Run(llb.Shlexf("curl -Lo /usr/bin/local-ai %s", localAIURL)).Root()
	c = c.Run(llb.Shlex("chmod +x /usr/bin/local-ai")).Root()

	diff := llb.Diff(initState, c)
	return llb.Merge([]llb.State{merge, diff})
}

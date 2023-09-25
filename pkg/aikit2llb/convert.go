package aikit2llb

import (
	"fmt"
	"net/url"
	"path"

	"github.com/moby/buildkit/client/llb"
	"github.com/sozercan/aikit/pkg/aikit/config"
)

const (
	debianSlim     = "docker.io/library/debian:12-slim"
	distrolessBase = "gcr.io/distroless/cc-debian12:latest"
	localAIVersion = "v1.40.0"
	retryCount     = 5
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
	merge := llb.Merge([]llb.State{llb.Image(distrolessBase), diff})
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
	i := s.Run(llb.Shlex("apt update"), llb.IgnoreCache).Root()
	return i.Run(llb.Shlex("apt install curl -y")).Root()
}

func addLocalAI(s llb.State, merge llb.State) llb.State {
	initState := s
	localAIURL := fmt.Sprintf("https://github.com/mudler/LocalAI/releases/download/%s/local-ai-avx-Linux-x86_64", localAIVersion)
	c := s.Run(llb.Shlexf("curl -Lo /usr/bin/local-ai %s", localAIURL)).Root()
	c = c.Run(llb.Shlex("chmod +x /usr/bin/local-ai")).Root()

	diff := llb.Diff(initState, c)
	return llb.Merge([]llb.State{merge, diff})
}

func shf(cmd string, v ...interface{}) llb.RunOption {
	return llb.Args([]string{"/bin/sh", "-c", fmt.Sprintf(cmd, v...)})
}

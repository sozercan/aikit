package inference

import (
	"github.com/moby/buildkit/client/llb"
	"github.com/sozercan/aikit/pkg/aikit/config"
	"github.com/sozercan/aikit/pkg/utils"
)

func installExllama(c *config.InferenceConfig, s llb.State, merge llb.State) llb.State {
	backend := utils.BackendExllama
	exllamaRepo := "https://github.com/turboderp/exllama"
	exllamaTag := "master"
	for b := range c.Backends {
		if c.Backends[b] == utils.BackendExllamaV2 {
			exllamaRepo = "https://github.com/turboderp/exllamav2"
			backend = utils.BackendExllamaV2
			exllamaTag = "v0.0.11"
			// c0ddebaaaf8ffd1b3529c2bb654e650bce2f790f
		}
	}

	savedState := s
	s = s.Run(utils.Sh("apt-get update && apt-get install --no-install-recommends -y git ca-certificates python3-pip python3-dev make g++ curl && curl -LsSf https://astral.sh/uv/install.sh | sh && source $HOME/.cargo/env && apt-get clean"), llb.IgnoreCache).Root()

	// clone localai exllama backend only
	s = cloneLocalAI(s, backend)

	// clone exllama to localai exllama backend path and install python dependencies
	s = s.Run(utils.Shf("git clone --depth 1 %[1]s --branch %[2]s /tmp/%[3]s && mv /tmp/%[3]s/* /tmp/localai/backend/python/%[3]s && rm -rf /tmp/%[3]s && cd /tmp/localai/backend/python/%[3]s && rm -rf .git && uv pip install --no-build-isolation --requirement requirements-install.txt && EXLLAMA_NOCOMPILE= uv pip install --no-build-isolation && make protogen", exllamaRepo, exllamaTag, backend)).Root()

	// && pip3 install grpcio protobuf typing-extensions sympy mpmath setuptools numpy --break-system-packages && pip3 install -r /tmp/localai/backend/python/%[3]s/requirements.txt --break-system-packages


	diff := llb.Diff(savedState, s)
	return llb.Merge([]llb.State{merge, diff})
}

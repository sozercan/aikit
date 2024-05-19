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
			exllamaTag = "v0.0.12"
		}
	}

	savedState := s
	s = s.Run(utils.Sh("apt-get update && apt-get install --no-install-recommends -y bash git ca-certificates python3-pip python3-dev python3-venv make g++ curl && curl -LsSf https://astral.sh/uv/install.sh | sh && pip install grpcio-tools --break-system-packages && apt-get clean"), llb.IgnoreCache).Root()

	// clone localai exllama backend only
	s = cloneLocalAI(s, backend)

	// clone exllama to localai exllama backend path and install python dependencies
	s = s.Run(utils.Bashf("git clone --depth 1 %[1]s --branch %[2]s /tmp/%[3]s && mv /tmp/%[3]s/* /tmp/localai/backend/python/%[3]s && rm -rf /tmp/%[3]s && cd /tmp/localai/backend/python/%[3]s && rm -rf .git && source $HOME/.cargo/env && python3 -m grpc_tools.protoc -I../.. --python_out=. --grpc_python_out=. backend.proto && uv venv && source .venv/bin/activate && ls -al && uv pip install --no-build-isolation --requirement requirements-install.txt && EXLLAMA_NOCOMPILE= uv pip install --no-build-isolation .", exllamaRepo, exllamaTag, backend)).Root()

	diff := llb.Diff(savedState, s)
	return llb.Merge([]llb.State{merge, diff})
}

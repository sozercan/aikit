package inference

import (
	"github.com/moby/buildkit/client/llb"
	"github.com/sozercan/aikit/pkg/aikit/config"
	"github.com/sozercan/aikit/pkg/utils"
)

func installExllama(c *config.InferenceConfig, s llb.State, merge llb.State) llb.State {
	backend := utils.BackendExllama
	for b := range c.Backends {
		if c.Backends[b] == utils.BackendExllamaV2 {
			backend = utils.BackendExllamaV2
		}
	}

	savedState := s
	s = s.Run(utils.Sh("apt-get update && apt-get install --no-install-recommends -y bash git ca-certificates python3-pip python3-dev python3-venv make g++ curl && curl -LsSf https://astral.sh/uv/install.sh | sh && pip install grpcio-tools --break-system-packages && apt-get clean"), llb.IgnoreCache).Root()

	// clone localai exllama backend only
	s = cloneLocalAI(s, backend)

	// clone exllama to localai exllama backend path and install python dependencies
	s = s.Run(utils.Shf("cd /tmp/localai/backend/python/%[1]s && rm -rf .git && . $HOME/.cargo/env && python3 -m grpc_tools.protoc -I../.. --python_out=. --grpc_python_out=. backend.proto && export BUILD_TYPE=cublas && ./install.sh", backend)).Root()

	diff := llb.Diff(savedState, s)
	return llb.Merge([]llb.State{merge, diff})
}

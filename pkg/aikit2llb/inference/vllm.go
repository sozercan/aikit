package inference

import (
	"github.com/moby/buildkit/client/llb"
	"github.com/sozercan/aikit/pkg/utils"
)

func installVLLM(s llb.State, merge llb.State) llb.State {
	savedState := s
	s = s.Run(utils.Sh("apt-get update && apt-get install --no-install-recommends -y bash git ca-certificates python3-pip python3-dev python3-venv python-is-python3 make g++ curl && pip install uv grpcio-tools && apt-get clean"), llb.IgnoreCache).Root()

	s = cloneLocalAI(s)

	s = s.Run(utils.Bashf("export BUILD_TYPE=cublas && export CUDA_MAJOR_VERSION=12 && cd /tmp/localai/backend/python/%[1]s && sed -i 's/grpcio==1.67.1/grpcio==1.68.0/g' requirements.txt && make %[1]s", utils.BackendVLLM)).Root()

	diff := llb.Diff(savedState, s)
	return llb.Merge([]llb.State{merge, diff})
}

package inference

import (
	"github.com/kaito-project/aikit/pkg/utils"
	"github.com/moby/buildkit/client/llb"
)

func installDiffusers(s llb.State, merge llb.State) llb.State {
	savedState := s
	s = s.Run(utils.Sh("apt-get install --no-install-recommends -y git python3 python3-pip python3-venv python-is-python3 make && pip install uv && pip install grpcio-tools==1.59.0 --no-dependencies && apt-get clean"), llb.IgnoreCache).Root()

	s = cloneLocalAI(s)

	s = s.Run(utils.Bashf("export BUILD_TYPE=cublas && export CUDA_MAJOR_VERSION=12 && cd /tmp/localai/backend/python/%[1]s && make %[1]s", utils.BackendDiffusers)).Root()

	diff := llb.Diff(savedState, s)
	return llb.Merge([]llb.State{merge, diff})
}

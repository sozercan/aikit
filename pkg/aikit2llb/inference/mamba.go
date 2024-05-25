package inference

import (
	"github.com/moby/buildkit/client/llb"
	"github.com/sozercan/aikit/pkg/utils"
)

func installMamba(s llb.State, merge llb.State) llb.State {
	savedState := s
	// libexpat1 is requirement but git is not. however libexpat1 is a dependency of git
	s = s.Run(utils.Sh("apt-get install --no-install-recommends -y git python3 python3-dev python3-pip python3-venv python-is-python3 libssl3 openssl curl && pip install uv grpcio-tools --break-system-packages && apt-get clean"), llb.IgnoreCache).Root()

	s = cloneLocalAI(s)

	s = s.Run(utils.Bashf("export BUILD_TYPE=cublas && cd /tmp/localai/backend/python/%[1]s && make %[1]s", utils.BackendMamba)).Root()

	diff := llb.Diff(savedState, s)
	return llb.Merge([]llb.State{merge, diff})
}

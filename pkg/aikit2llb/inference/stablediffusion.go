package inference

import (
	"github.com/moby/buildkit/client/llb"
	"github.com/sozercan/aikit/pkg/utils"
)

func installOpenCV(s llb.State, merge llb.State) llb.State {
	savedState := s
	libPath := "/usr/lib/x86_64-linux-gnu"
	// add symbolic opencv 4.5d links for stablediffusion on bookworm
	s = s.Run(utils.Shf("apt-get update && mkdir -p /tmp/generated/images && apt-get install --no-install-recommends -y curl unzip ca-certificates libopencv-imgcodecs406 libgomp1 libdap27 libdapclient6v5 && apt-get clean && ln -s %[1]s/libopencv_core.so.4.6.0 %[1]s/libopencv_core.so.4.5d && ln -s %[1]s/libopencv_imgcodecs.so.4.6.0 %[1]s/libopencv_imgcodecs.so.4.5d", libPath), llb.IgnoreCache).Root()
	diff := llb.Diff(savedState, s)
	merge = llb.Merge([]llb.State{merge, diff})

	// https://github.com/mudler/LocalAI/actions/runs/9227834555 (v2.16.0)
	// temporary fix for stablediffusion
	sdURL := "https://nightly.link/mudler/LocalAI/actions/runs/9227834555/stablediffusion.zip"
	merge = merge.Run(utils.Shf("mkdir -p /tmp/localai/backend_data/backend-assets/grpc/ && curl --retry 10 --retry-all-errors -L %s -o stablediffusion.zip && unzip stablediffusion.zip -d /tmp/localai/backend_data/backend-assets/grpc && chmod +x /tmp/localai/backend_data/backend-assets/grpc/stablediffusion", sdURL)).Root()
	return merge
}

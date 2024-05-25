package inference

import (
	"github.com/moby/buildkit/client/llb"
	"github.com/sozercan/aikit/pkg/utils"
)

func installOpenCV(s llb.State, merge llb.State) llb.State {
	savedState := s
	// adding debian 11 (bullseye) repo due to opencv 4.5 requirement
	s = s.Run(utils.Sh("echo 'deb http://deb.debian.org/debian bullseye main' | tee -a /etc/apt/sources.list")).Root()
	// pinning libdap packages to bullseye version due to symbol error
	libdapVersion := "3.20.7-6"
	libPath := "/usr/lib/x86_64-linux-gnu"
	s = s.Run(utils.Shf("apt-get update && mkdir -p /tmp/generated/images && apt-get install --no-install-recommends -y curl unzip ca-certificates libopencv-imgcodecs4.5 libgomp1 libdap27=%[1]s libdapclient6v5=%[1]s && apt-get clean && ln -s %[2]s/libopencv_core.so.4.5 %[2]s/libopencv_core.so.4.5d && ln -s %[2]s/libopencv_core.so.4.5 %[2]s/libopencv_core.so.406 && ln -s %[2]s/libopencv_imgcodecs.so.4.5 %[2]s/libopencv_imgcodecs.so.4.5d", libdapVersion, libPath), llb.IgnoreCache).Root()
	diff := llb.Diff(savedState, s)
	merge = llb.Merge([]llb.State{merge, diff})

	// https://github.com/mudler/LocalAI/actions/runs/9227834555 (v2.16.0)
	// temporary fix for stablediffusion
	sdURL := "https://nightly.link/mudler/LocalAI/actions/runs/9227834555/stablediffusion.zip"
	merge = merge.Run(utils.Shf("mkdir -p /tmp/localai/backend_data/backend-assets/grpc/ && curl --retry 10 --retry-all-errors -L %s -o stablediffusion.zip && unzip stablediffusion.zip -d /tmp/localai/backend_data/backend-assets/grpc && chmod +x stablediffusion", sdURL)).Root()
	return merge
}

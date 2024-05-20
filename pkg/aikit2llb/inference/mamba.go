package inference

import (
	"fmt"

	"github.com/moby/buildkit/client/llb"
	"github.com/sozercan/aikit/pkg/utils"
)

func installMamba(s llb.State, merge llb.State) llb.State {
	savedState := s
	// libexpat1 is requirement but git is not. however libexpat1 is a dependency of git
	s = s.Run(utils.Sh("apt-get install --no-install-recommends -y git python3 python3-dev python3-pip libssl3 openssl && apt-get clean"), llb.IgnoreCache).Root()

	s = cloneLocalAI(s)

	s = s.Run(utils.Shf("pip3 install packaging numpy torch==2.1.0 grpcio protobuf --break-system-packages && pip3 install causal-conv1d==1.0.0 mamba-ssm==1.0.1 --break-system-packages")).Root()

	diff := llb.Diff(savedState, s)
	return llb.Merge([]llb.State{merge, diff})
}

func installOpenCV(s llb.State, merge llb.State) llb.State {
	savedState := s
	// adding debian 11 (bullseye) repo due to opencv 4.5 requirement
	s = s.Run(utils.Sh("echo 'deb http://deb.debian.org/debian bullseye main' | tee -a /etc/apt/sources.list")).Root()
	// pinning libdap packages to bullseye version due to symbol error
	libdapVersion := "3.20.7-6"
	libPath := "/usr/lib/x86_64-linux-gnu"
	s = s.Run(utils.Shf("apt-get update && mkdir -p /tmp/generated/images && apt-get install -y libopencv-imgcodecs4.5 libgomp1 libdap27=%[1]s libdapclient6v5=%[1]s && apt-get clean && ln -s %[2]s/libopencv_core.so.4.5 %[2]s/libopencv_core.so.4.5d && ln -s %[2]s/libopencv_imgcodecs.so.4.5 %[2]s/libopencv_imgcodecs.so.4.5d", libdapVersion, libPath), llb.IgnoreCache).Root()
	diff := llb.Diff(savedState, s)
	merge = llb.Merge([]llb.State{merge, diff})

	sdURL := fmt.Sprintf("https://sertaccdn.azureedge.net/localai/%s/stablediffusion", localAIVersion)
	var opts []llb.HTTPOption
	opts = append(opts, llb.Filename("stablediffusion"))
	opts = append(opts, llb.Chmod(0o755))
	var copyOpts []llb.CopyOption
	copyOpts = append(copyOpts, &llb.CopyInfo{
		CreateDestPath: true,
	})
	sd := llb.HTTP(sdURL, opts...)
	merge = merge.File(
		llb.Copy(sd, "stablediffusion", "/tmp/localai/backend_data/backend-assets/grpc/stablediffusion", copyOpts...),
		llb.WithCustomName("Copying stable diffusion backend"), //nolint: goconst
	)
	return merge
}

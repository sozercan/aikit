package inference

import (
	"fmt"

	"github.com/moby/buildkit/client/llb"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sozercan/aikit/pkg/utils"
)

func installOpenCV(s llb.State, merge llb.State, platform specs.Platform) llb.State {
	// pinning libdap packages to bullseye version due to symbol error
	libdapVersion := "3.20.7-6"
	libPaths := map[string]string{
		utils.PlatformAMD64: "/usr/lib/x86_64-linux-gnu",
		utils.PlatformARM64: "/usr/lib/aarch64-linux-gnu",
	}
	libPath, exists := libPaths[platform.Architecture]
	if !exists {
		return s
	}

	savedState := s
	// adding debian 11 (bullseye) repo due to opencv 4.5 requirement
	s = s.Run(utils.Sh("echo 'deb http://deb.debian.org/debian bullseye main' | tee -a /etc/apt/sources.list")).Root()
	s = s.Run(utils.Shf("apt-get update && mkdir -p /tmp/generated/images && apt-get install --no-install-recommends -y curl unzip ca-certificates libopencv-imgcodecs4.5 libgomp1 libdap27=%[1]s libdapclient6v5=%[1]s && apt-get clean && ln -s %[2]s/libopencv_core.so.4.5 %[2]s/libopencv_core.so.4.5d && ln -s %[2]s/libopencv_core.so.4.5 %[2]s/libopencv_core.so.406 && ln -s %[2]s/libopencv_imgcodecs.so.4.5 %[2]s/libopencv_imgcodecs.so.4.5d", libdapVersion, libPath), llb.IgnoreCache).Root()
	diff := llb.Diff(savedState, s)
	merge = llb.Merge([]llb.State{merge, diff})

	// TODO: update this URL when the binary is available in github
	sdURL := fmt.Sprintf("https://sertaccdnvs.azureedge.net/localai/%s/stablediffusion", localAIVersion)
	var opts []llb.HTTPOption
	opts = append(opts, llb.Filename("stablediffusion"), llb.Chmod(0o755))
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

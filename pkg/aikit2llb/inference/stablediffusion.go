package inference

import (
	"fmt"

	"github.com/moby/buildkit/client/llb"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sozercan/aikit/pkg/utils"
)

func installOpenCV(s llb.State, merge llb.State, platform specs.Platform) llb.State {
	libPaths := map[string]string{
		utils.PlatformAMD64: "/usr/lib/x86_64-linux-gnu",
		utils.PlatformARM64: "/usr/lib/aarch64-linux-gnu",
	}
	libPath, exists := libPaths[platform.Architecture]
	if !exists {
		return s
	}

	savedState := s

	s = s.Run(utils.Shf("apt-get update && mkdir -p /tmp/generated/images && apt-get install --no-install-recommends -y libopencv-imgcodecs4.5d && ln -s %[1]s/libopencv_core.so.4.5d %[1]s/libopencv_core.so.406 && ln -s %[1]s/libopencv_imgcodecs.so.4.5d %[1]s/libopencv_imgcodecs.so.406 && ln -s %[1]s/libopencv_imgproc.so.4.5d %[1]s/libopencv_imgproc.so.406 && apt-get clean", libPath), llb.IgnoreCache).Root()
	diff := llb.Diff(savedState, s)
	merge = llb.Merge([]llb.State{merge, diff})

	sdURL := fmt.Sprintf("https://github.com/mudler/LocalAI/releases/download/%[1]s/stablediffusion", localAIVersion)
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

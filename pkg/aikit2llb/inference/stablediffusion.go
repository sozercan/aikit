package inference

import (
	"fmt"

	"github.com/moby/buildkit/client/llb"
	"github.com/sozercan/aikit/pkg/utils"
)

func installOpenCV(s llb.State, merge llb.State) llb.State {
	savedState := s

	s = s.Run(utils.Sh("apt-get update && mkdir -p /tmp/generated/images && apt-get install --no-install-recommends -y libopencv-imgcodecs4.5d && apt-get clean"), llb.IgnoreCache).Root()
	diff := llb.Diff(savedState, s)
	merge = llb.Merge([]llb.State{merge, diff})

	sdURL := fmt.Sprintf("https://sertaccdnvs.azureedge.net/localai/%[1]s/stablediffusion/stablediffusion", localAIVersion)
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

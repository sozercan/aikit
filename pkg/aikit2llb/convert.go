package aikit2llb

import (
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/moby/buildkit/client/llb"
	"github.com/opencontainers/go-digest"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/sozercan/aikit/pkg/aikit/config"
	"github.com/sozercan/aikit/pkg/utils"
)

const (
	debianSlim     = "docker.io/library/debian:12-slim"
	distrolessBase = "gcr.io/distroless/cc-debian12:latest"
	localAIVersion = "v2.0.0"
	cudaVersion    = "12-3"
)

func Aikit2LLB(c *config.Config) (llb.State, *specs.Image) {
	var merge llb.State
	s := llb.Image(debianSlim)
	s, merge = copyModels(c, s)
	s, merge = addLocalAI(c, s, merge)
	if c.Runtime == utils.RuntimeNVIDIA {
		merge = installCuda(s, merge)
	}
	for b := range c.Backends {
		if strings.Contains(c.Backends[b], "stablediffusion") {
			merge = installOpenCV(s, merge)
		}
	}
	imageCfg := NewImageConfig(c)
	return merge, imageCfg
}

func copyModels(c *config.Config, s llb.State) (llb.State, llb.State) {
	var base llb.State
	for b := range c.Backends {
		if strings.Contains(c.Backends[b], "stablediffusion") {
			// due to too many dependencies, we are using debian slim as base for stable diffusion
			base = llb.Image(debianSlim)
		} else {
			base = llb.Image(distrolessBase)
		}
	}
	savedState := s

	// create config file if defined
	if c.Config != "" {
		s = s.Run(shf("echo \"%s\" >> /config.yaml", c.Config)).Root()
	}

	for _, model := range c.Models {
		var opts []llb.HTTPOption
		opts = append(opts, llb.Filename(fileNameFromURL(model.Source)))
		if model.SHA256 != "" {
			digest := digest.NewDigestFromEncoded(digest.SHA256, model.SHA256)
			opts = append(opts, llb.Checksum(digest))
		}

		m := llb.HTTP(model.Source, opts...)

		var modelPath string
		if strings.Contains(model.Name, "/") {
			modelPath = "/models/" + path.Dir(model.Name) + "/" + fileNameFromURL(model.Source)
		} else {
			modelPath = "/models/" + fileNameFromURL(model.Source)
		}

		var copyOpts []llb.CopyOption
		copyOpts = append(copyOpts, &llb.CopyInfo{
			CreateDestPath: true,
		})
		s = s.File(
			llb.Copy(m, fileNameFromURL(model.Source), modelPath, copyOpts...),
			llb.WithCustomName("Copying "+fileNameFromURL(model.Source)+" to "+modelPath), //nolint: goconst
		)

		// create prompt templates if defined
		for _, pt := range model.PromptTemplates {
			if pt.Name != "" && pt.Template != "" {
				s = s.Run(shf("echo \"%s\" >> /models/%s.tmpl", pt.Template, pt.Name)).Root()
			}
		}
	}
	diff := llb.Diff(savedState, s)
	merge := llb.Merge([]llb.State{base, diff})
	return s, merge
}

func fileNameFromURL(urlString string) string {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}
	return path.Base(parsedURL.Path)
}

func installCuda(s llb.State, merge llb.State) llb.State {
	cudaKeyringURL := "https://developer.download.nvidia.com/compute/cuda/repos/debian12/x86_64/cuda-keyring_1.1-1_all.deb"
	cudaKeyring := llb.HTTP(cudaKeyringURL)
	s = s.File(
		llb.Copy(cudaKeyring, fileNameFromURL(cudaKeyringURL), "/"),
		llb.WithCustomName("Copying "+fileNameFromURL(cudaKeyringURL)), //nolint: goconst
	)
	s = s.Run(shf("dpkg -i cuda-keyring_1.1-1_all.deb && rm cuda-keyring_1.1-1_all.deb")).Root()
	// running apt-get update twice due to nvidia repo
	s = s.Run(shf("apt-get update && apt-get install -y ca-certificates && apt-get update"), llb.IgnoreCache).Root()
	savedState := s
	s = s.Run(shf("apt-get install -y libcublas-%[1]s cuda-cudart-%[1]s && apt-get clean", cudaVersion)).Root()

	diff := llb.Diff(savedState, s)
	merge = llb.Merge([]llb.State{merge, diff})
	return merge
}

func installOpenCV(s llb.State, merge llb.State) llb.State {
	savedState := s
	// adding debian 11 (bullseye) repo due to opencv 4.5 requirement
	s = s.Run(shf("echo 'deb http://deb.debian.org/debian bullseye main' | tee -a /etc/apt/sources.list")).Root()
	// pinning libdap packages to bullseye version due to symbol error
	s = s.Run(shf("apt-get update && mkdir -p /tmp/generated/images && apt-get install -y libopencv-imgcodecs4.5 libgomp1 libdap27=3.20.7-6 libdapclient6v5=3.20.7-6 && apt-get clean"), llb.IgnoreCache).Root()
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

func addLocalAI(c *config.Config, s llb.State, merge llb.State) (llb.State, llb.State) {
	savedState := s
	var localAIURL string
	switch c.Runtime {
	case utils.RuntimeNVIDIA:
		localAIURL = fmt.Sprintf("https://sertaccdn.azureedge.net/localai/%s/cuda12/local-ai", localAIVersion)
	case utils.RuntimeCPUAVX2:
		localAIURL = fmt.Sprintf("https://github.com/mudler/LocalAI/releases/download/%s/local-ai-avx2-Linux-x86_64", localAIVersion)
	case utils.RuntimeCPUAVX512:
		localAIURL = fmt.Sprintf("https://github.com/mudler/LocalAI/releases/download/%s/local-ai-avx512-Linux-x86_64", localAIVersion)
	case utils.RuntimeCPUAVX, "":
		localAIURL = fmt.Sprintf("https://github.com/mudler/LocalAI/releases/download/%s/local-ai-avx-Linux-x86_64", localAIVersion)
	}

	var opts []llb.HTTPOption
	opts = append(opts, llb.Filename("local-ai"))
	opts = append(opts, llb.Chmod(0o755))
	localAI := llb.HTTP(localAIURL, opts...)
	s = s.File(
		llb.Copy(localAI, "local-ai", "/usr/bin"),
		llb.WithCustomName("Copying "+fileNameFromURL(localAIURL)+" to /usr/bin"), //nolint: goconst
	)

	diff := llb.Diff(savedState, s)
	return s, llb.Merge([]llb.State{merge, diff})
}

func shf(cmd string, v ...interface{}) llb.RunOption {
	return llb.Args([]string{"/bin/sh", "-c", fmt.Sprintf(cmd, v...)})
}

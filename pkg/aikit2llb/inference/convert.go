package inference

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
	distrolessBase = "gcr.io/distroless/cc-debian12:latest"

	localAIRepo    = "https://github.com/mudler/LocalAI"
	localAIVersion = "v2.16.0"
	cudaVersion    = "12-3"
)

// Aikit2LLB converts an InferenceConfig to an LLB state.
func Aikit2LLB(c *config.InferenceConfig, platform *specs.Platform) (llb.State, *specs.Image, error) {
	var merge llb.State
	state := llb.Image(utils.DebianSlim, llb.Platform(*platform))
	base := getBaseImage(c, platform)

	state, merge = copyModels(c, base, state, *platform)

	var err error
	state, merge, err = addLocalAI(state, merge, *platform)
	if err != nil {
		return state, nil, err
	}

	// install cuda if runtime is nvidia and architecture is amd64
	if c.Runtime == utils.RuntimeNVIDIA && platform.Architecture == utils.PlatformAMD64 {
		state, merge = installCuda(c, state, merge)
	}

	// install backend dependencies
	for b := range c.Backends {
		switch c.Backends[b] {
		case utils.BackendExllama, utils.BackendExllamaV2:
			merge = installExllama(c, state, merge)
		case utils.BackendStableDiffusion:
			merge = installOpenCV(state, merge, *platform)
		case utils.BackendMamba:
			merge = installMamba(state, merge)
		}
	}

	imageCfg := NewImageConfig(c, platform)
	return merge, imageCfg, nil
}

// getBaseImage returns the base image given the InferenceConfig and platform.
func getBaseImage(c *config.InferenceConfig, platform *specs.Platform) llb.State {
	if len(c.Backends) > 0 {
		return llb.Image(utils.DebianSlim, llb.Platform(*platform))
	}
	return llb.Image(distrolessBase, llb.Platform(*platform))
}

// copyModels copies models to the image.
func copyModels(c *config.InferenceConfig, base llb.State, s llb.State, platform specs.Platform) (llb.State, llb.State) {
	savedState := s
	for _, model := range c.Models {
		// check if model source is a URL or a local path
		_, err := url.ParseRequestURI(model.Source)
		if err == nil {
			// download from oci artifacts
			if strings.Contains(model.Source, "oci://") {
				// TODO: replace this
				craneBase := "docker.io/alpine/crane:latest"
				toolingImage := llb.Image(craneBase, llb.Platform(platform))

				artifactURL := strings.TrimPrefix(model.Source, "oci://")
				const ollamaRegistryURL = "registry.ollama.ai"
				var craneCmd, modelName string
				if strings.HasPrefix(artifactURL, ollamaRegistryURL) {
					// remove the tag so we can append the digest
					artifactURLWithoutTag := strings.Split(artifactURL, ":")[0]
					// extract name of the model from registry.ollama.ai/namespace/name
					modelName := strings.Split(artifactURLWithoutTag, "/")[2] + ".gguf"
					// model is stored with media type application/vnd.ollama.image.model
					craneCmd = fmt.Sprintf("crane blob %[1]s@$(crane manifest %[2]s | jq -r '.layers[] | select(.mediaType == \"application/vnd.ollama.image.model\").digest') > %[3]s", artifactURLWithoutTag, artifactURL, modelName)
				} else {
					// generic oci artifact
					modelName := path.Base(artifactURL)
					if strings.Contains(modelName, ":") {
						modelName = strings.Split(modelName, ":")[0]
					}
					if strings.Contains(modelName, "@") {
						modelName = strings.Split(modelName, "@")[0]
					}
					craneCmd = fmt.Sprintf("crane blob %[1]s > /models/%[2]s", artifactURL, modelName)
				}

				toolingImage = toolingImage.Run(utils.Sh("apk add jq")).Root()
				toolingImage = toolingImage.Run(utils.Sh(craneCmd)).Root()

				var copyOpts []llb.CopyOption
				copyOpts = append(copyOpts, &llb.CopyInfo{
					CreateDestPath: true,
				})
				modelPath := fmt.Sprintf("/models/%s", modelName)
				s = toolingImage.File(
					llb.Copy(toolingImage, modelName, modelPath, copyOpts...),
					llb.WithCustomName("Copying "+artifactURL+" to "+modelPath), //nolint: goconst
				)
			} else {
				// http download
				var opts []llb.HTTPOption
				opts = append(opts, llb.Filename(utils.FileNameFromURL(model.Source)))
				if model.SHA256 != "" {
					digest := digest.NewDigestFromEncoded(digest.SHA256, model.SHA256)
					opts = append(opts, llb.Checksum(digest))
				}

				m := llb.HTTP(model.Source, opts...)

				var modelPath string
				if strings.Contains(model.Name, "/") {
					modelPath = "/models/" + path.Dir(model.Name) + "/" + utils.FileNameFromURL(model.Source)
				} else {
					modelPath = "/models/" + utils.FileNameFromURL(model.Source)
				}

				var copyOpts []llb.CopyOption
				copyOpts = append(copyOpts, &llb.CopyInfo{
					CreateDestPath: true,
				})
				s = s.File(
					llb.Copy(m, utils.FileNameFromURL(model.Source), modelPath, copyOpts...),
					llb.WithCustomName("Copying "+utils.FileNameFromURL(model.Source)+" to "+modelPath), //nolint: goconst
				)
			}
		} else {
			// copy from local path
			var copyOpts []llb.CopyOption
			copyOpts = append(copyOpts, &llb.CopyInfo{
				CreateDestPath: true,
			})
			s = s.File(
				llb.Copy(llb.Local("context"), model.Source, "/models/", copyOpts...),
				llb.WithCustomName("Copying "+utils.FileNameFromURL(model.Source)+" to "+"/models"), //nolint: goconst
			)
		}

		// create prompt templates if defined
		for _, pt := range model.PromptTemplates {
			if pt.Name != "" && pt.Template != "" {
				s = s.Run(utils.Shf("echo -n \"%s\" > /models/%s.tmpl", pt.Template, pt.Name)).Root()
			}
		}
	}

	// create config file if defined
	if c.Config != "" {
		s = s.Run(utils.Shf("mkdir -p /configuration && echo -n \"%s\" > /config.yaml", c.Config)).Root()
	}

	diff := llb.Diff(savedState, s)
	merge := llb.Merge([]llb.State{base, diff})
	return s, merge
}

// installCuda installs cuda libraries and dependencies.
func installCuda(c *config.InferenceConfig, s llb.State, merge llb.State) (llb.State, llb.State) {
	cudaKeyringURL := "https://developer.download.nvidia.com/compute/cuda/repos/debian12/x86_64/cuda-keyring_1.1-1_all.deb"
	cudaKeyring := llb.HTTP(cudaKeyringURL)
	s = s.File(
		llb.Copy(cudaKeyring, utils.FileNameFromURL(cudaKeyringURL), "/"),
		llb.WithCustomName("Copying "+utils.FileNameFromURL(cudaKeyringURL)), //nolint: goconst
	)
	s = s.Run(utils.Sh("dpkg -i cuda-keyring_1.1-1_all.deb && rm cuda-keyring_1.1-1_all.deb")).Root()

	savedState := s
	// running apt-get update twice due to nvidia repo
	s = s.Run(utils.Sh("apt-get update && apt-get install --no-install-recommends -y ca-certificates && apt-get update"), llb.IgnoreCache).Root()

	// default llama.cpp backend is being used
	if len(c.Backends) == 0 {
		// install cuda libraries and pciutils for gpu detection
		s = s.Run(utils.Shf("apt-get install -y --no-install-recommends pciutils libcublas-%[1]s cuda-cudart-%[1]s && apt-get clean", cudaVersion)).Root()
		// using a distroless base image here
		// convert debian package metadata status file to distroless status.d directory
		// clean up apt directories
		s = s.Run(utils.Bashf("apt-get install -y --no-install-recommends libcublas-%[1]s cuda-cudart-%[1]s && apt-get clean && mkdir -p /var/lib/dpkg/status.d && description_flag=false; while IFS= read -r line || [[ -n $line ]]; do if [[ $line == Package:* ]]; then pkg_name=$(echo $line | cut -d' ' -f2); elif [[ $line == Maintainer:* ]]; then maintainer=$(echo $line | cut -d' ' -f2-); if [[ $maintainer == 'cudatools <cudatools@nvidia.com>' ]]; then pkg_file=/var/lib/dpkg/status.d/${pkg_name}; echo 'Package: '$pkg_name > $pkg_file; echo $line >> $pkg_file; else pkg_file=''; fi; elif [[ -n $pkg_file ]]; then if [[ $line == Description:* ]]; then description_flag=true; elif [[ $line == '' ]]; then description_flag=false; elif ! $description_flag; then echo $line >> $pkg_file; fi; fi; done < /var/lib/dpkg/status && find /var/lib/dpkg -mindepth 1 ! -regex '^/var/lib/dpkg/status\\.d\\(/.*\\)?' -delete && rm -r /var/lib/apt", cudaVersion)).Root()
	}

	// installing dev dependencies used for exllama
	for b := range c.Backends {
		if c.Backends[b] == utils.BackendExllama || c.Backends[b] == utils.BackendExllamaV2 {
			var exllama2Dep string
			if c.Backends[b] == utils.BackendExllamaV2 {
				exllama2Dep = fmt.Sprintf("libcurand-dev-%[1]s", cudaVersion)
			}
			exllamaDeps := fmt.Sprintf("apt-get install -y --no-install-recommends cuda-cudart-dev-%[1]s cuda-crt-%[1]s libcusparse-dev-%[1]s libcublas-dev-%[1]s libcusolver-dev-%[1]s cuda-nvcc-%[1]s %[2]s && apt-get clean", cudaVersion, exllama2Dep)

			s = s.Run(utils.Sh(exllamaDeps)).Root()
		}

		if c.Backends[b] == utils.BackendMamba {
			mambaDeps := fmt.Sprintf("apt-get install -y --no-install-recommends cuda-crt-%[1]s cuda-cudart-dev-%[1]s cuda-nvcc-%[1]s && apt-get clean", cudaVersion)
			s = s.Run(utils.Sh(mambaDeps)).Root()
		}
	}

	diff := llb.Diff(savedState, s)
	return s, llb.Merge([]llb.State{merge, diff})
}

// addLocalAI adds the LocalAI binary to the image.
func addLocalAI(s llb.State, merge llb.State, platform specs.Platform) (llb.State, llb.State, error) {
	binaryNames := map[string]string{
		utils.PlatformAMD64: "local-ai-Linux-x86_64",
		utils.PlatformARM64: "local-ai-Linux-arm64",
	}
	binaryName, exists := binaryNames[platform.Architecture]
	if !exists {
		return s, merge, fmt.Errorf("unsupported architecture %s", platform.Architecture)
	}
	// TODO: update this URL when the binary is available in github
	localAIURL := fmt.Sprintf("https://sertacstoragevs.blob.core.windows.net/localai/%[1]s/%[2]s", localAIVersion, binaryName)

	savedState := s

	var opts []llb.HTTPOption
	opts = append(opts, llb.Filename("local-ai"), llb.Chmod(0o755))
	localAI := llb.HTTP(localAIURL, opts...)
	s = s.File(
		llb.Copy(localAI, "local-ai", "/usr/bin/local-ai"),
		llb.WithCustomName("Copying "+utils.FileNameFromURL(localAIURL)+" to /usr/bin"), //nolint: goconst
	)

	diff := llb.Diff(savedState, s)
	return s, llb.Merge([]llb.State{merge, diff}), nil
}

// cloneLocalAI clones the LocalAI repository to the image used for python backends.
func cloneLocalAI(s llb.State) llb.State {
	return s.Run(utils.Shf("git clone --filter=blob:none --no-checkout %[1]s /tmp/localai/ && cd /tmp/localai && git sparse-checkout init --cone && git sparse-checkout set backend/python && git checkout %[2]s && rm -rf .git", localAIRepo, localAIVersion)).Root()
}

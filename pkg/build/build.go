package build

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/containerd/containerd/platforms"
	"github.com/moby/buildkit/client/llb"
	"github.com/moby/buildkit/exporter/containerimage/exptypes"
	"github.com/moby/buildkit/frontend/dockerui"
	"github.com/moby/buildkit/frontend/gateway/client"
	"github.com/moby/buildkit/util/gitutil"
	"github.com/pkg/errors"
	"github.com/sozercan/aikit/pkg/aikit/config"
	"github.com/sozercan/aikit/pkg/aikit2llb/finetune"
	"github.com/sozercan/aikit/pkg/aikit2llb/inference"
	"github.com/sozercan/aikit/pkg/utils"
)

const (
	localNameContext     = "context"
	localNameDockerfile  = "dockerfile"
	localNameAikitfile   = "aikitfile.yaml"
	defaultAikitfileName = "aikitfile.yaml"
	keyFilename          = "filename"
	keyTarget            = "target"
	keyOutput            = "output"
)

func Build(ctx context.Context, c client.Client) (*client.Result, error) {
	inferenceCfg, finetuneCfg, err := getAikitfileConfig(ctx, c)
	if err != nil {
		return nil, errors.Wrap(err, "getting aikitfile")
	}

	if finetuneCfg != nil {
		return buildFineTune(ctx, c, finetuneCfg)
	} else if inferenceCfg != nil {
		return buildInference(ctx, c, inferenceCfg)
	}

	return nil, nil
}

func buildFineTune(ctx context.Context, c client.Client, cfg *config.FineTuneConfig) (*client.Result, error) {
	err := validateFinetuneConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "validating aikitfile")
	}

	// set defaults for unsloth and finetune config
	if cfg.Target == utils.TargetUnsloth {
		cfg = defaultsUnslothConfig(cfg)
	}
	cfg = defaultsFineTune(cfg)

	st := finetune.Aikit2LLB(cfg)

	def, err := st.Marshal(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal local source")
	}
	res, err := c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve dockerfile")
	}
	return res, nil
}

func buildInference(ctx context.Context, c client.Client, cfg *config.InferenceConfig) (*client.Result, error) {
	err := validateInferenceConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "validating aikitfile")
	}

	st, img := inference.Aikit2LLB(cfg)

	def, err := st.Marshal(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal local source")
	}
	res, err := c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to resolve dockerfile")
	}
	ref, err := res.SingleRef()
	if err != nil {
		return nil, err
	}

	config, err := json.Marshal(img)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to marshal image config")
	}
	k := platforms.Format(platforms.DefaultSpec())

	res.AddMeta(fmt.Sprintf("%s/%s", exptypes.ExporterImageConfigKey, k), config)
	res.SetRef(ref)

	return res, nil
}

func getAikitfileConfig(ctx context.Context, c client.Client) (*config.InferenceConfig, *config.FineTuneConfig, error) {
	opts := c.BuildOpts().Opts
	filename := opts[keyFilename]
	if filename == "" {
		filename = defaultAikitfileName
	}

	name := "load aikitfile"
	if filename != "aikitfile.yaml" {
		name += " from " + filename
	}

	context := opts[localNameContext]

	var st llb.State
	var err error
	switch {
	case strings.HasPrefix(context, "git"):
		st, err = DetectGitContext(context, true)
		if err != nil {
			return nil, nil, err
		}
	case strings.HasPrefix(context, "http") || strings.HasPrefix(context, "https"):
		st, err = DetectGitContext(context, true)
		if err != nil {
			st = llb.HTTP(context, llb.WithCustomName("[context] "+context))
		}
	default:
		st = llb.Local(localNameDockerfile,
			llb.IncludePatterns([]string{filename}),
			llb.SessionID(c.BuildOpts().SessionID),
			llb.SharedKeyHint(defaultAikitfileName),
			dockerui.WithInternalName(name),
		)
	}

	def, err := st.Marshal(ctx)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to marshal local source")
	}

	res, err := c.Solve(ctx, client.SolveRequest{
		Definition: def.ToPB(),
	})
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to resolve aikitfile")
	}

	ref, err := res.SingleRef()
	if err != nil {
		return nil, nil, err
	}

	dtAikitfile, err := ref.ReadFile(ctx, client.ReadRequest{
		Filename: filename,
	})
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to read aikitfile")
	}

	inferenceCfg, finetuneCfg, err := config.NewFromBytes(dtAikitfile)
	if err != nil {
		return nil, nil, errors.Wrap(err, "getting config")
	}
	if finetuneCfg != nil {
		target, ok := opts[keyTarget]
		if !ok {
			target = utils.TargetUnsloth
		}
		finetuneCfg.Target = target

		if opts[keyOutput] != "" {
			return nil, nil, errors.New("--output is required for finetune. please specify a directory to save the finetuned model")
		}
	}

	// parse build args
	if inferenceCfg != nil {
		modelID := getBuildArg(opts, "MODEL_ID")
		modelFile := getBuildArg(opts, "MODEL_FILE")
		if modelID != "" && modelFile != "" {
			if !strings.HasPrefix(modelID, "huggingface://") {
				return nil, nil, errors.New("only huggingface models are supported at this time")
			}
			if !strings.HasSuffix(modelFile, ".gguf") {
				return nil, nil, errors.New("only GGUF files are supported at this time")
			}

			inferenceCfg.Models[0].Name = modelFile
			modelID = strings.TrimPrefix(modelID, "huggingface://")
			inferenceCfg.Models[0].Source = "https://huggingface.co/" + modelID + "/resolve/main/" + modelFile
			inferenceCfg.Config = fmt.Sprintf(`
- name: %[1]s
  backend: llama
  parameters:
    model: %[1]s`, modelFile)
		}
	}

	return inferenceCfg, finetuneCfg, nil
}

func getBuildArg(opts map[string]string, k string) string {
	if opts != nil {
		if v, ok := opts["build-arg:"+k]; ok {
			return v
		}
	}
	return ""
}

func validateFinetuneConfig(c *config.FineTuneConfig) error {
	supportedFineTuneTargets := []string{utils.TargetUnsloth}

	if c.APIVersion == "" {
		return errors.New("apiVersion is not defined")
	}

	if c.APIVersion != utils.APIv1alpha1 {
		return errors.Errorf("apiVersion %s is not supported", c.APIVersion)
	}

	if !slices.Contains(supportedFineTuneTargets, c.Target) {
		return errors.Errorf("target %s is not supported", c.Target)
	}

	if len(c.Datasets) == 0 {
		return errors.New("no datasets defined")
	}

	if len(c.Datasets) > 1 {
		return errors.New("only one dataset is supported at this time")
	}

	// only alpaca dataset is supported at this time
	for _, d := range c.Datasets {
		if d.Type != utils.DatasetAlpaca {
			return errors.Errorf("dataset type %s is not supported", d.Type)
		}
	}
	return nil
}

func defaultsUnslothConfig(c *config.FineTuneConfig) *config.FineTuneConfig {
	if c.Config.Unsloth.MaxSeqLength == 0 {
		c.Config.Unsloth.MaxSeqLength = 2048
	}
	if c.Config.Unsloth.BatchSize == 0 {
		c.Config.Unsloth.BatchSize = 2
	}
	if c.Config.Unsloth.GradientAccumulationSteps == 0 {
		c.Config.Unsloth.GradientAccumulationSteps = 4
	}
	if c.Config.Unsloth.WarmupSteps == 0 {
		c.Config.Unsloth.WarmupSteps = 10
	}
	if c.Config.Unsloth.MaxSteps == 0 {
		c.Config.Unsloth.MaxSteps = 60
	}
	if c.Config.Unsloth.LearningRate == 0 {
		c.Config.Unsloth.LearningRate = 0.0002
	}
	if c.Config.Unsloth.LoggingSteps == 0 {
		c.Config.Unsloth.LoggingSteps = 1
	}
	if c.Config.Unsloth.Optimizer == "" {
		c.Config.Unsloth.Optimizer = "adamw_8bit"
	}
	if c.Config.Unsloth.WeightDecay == 0 {
		c.Config.Unsloth.WeightDecay = 0.01
	}
	if c.Config.Unsloth.LrSchedulerType == "" {
		c.Config.Unsloth.LrSchedulerType = "linear"
	}
	if c.Config.Unsloth.Seed == 0 {
		c.Config.Unsloth.Seed = 42
	}
	return c
}

func defaultsFineTune(c *config.FineTuneConfig) *config.FineTuneConfig {
	if c.Output.Quantize == "" {
		c.Output.Quantize = "q4_k_m"
	}
	if c.Output.Name == "" {
		c.Output.Name = "aikit-model"
	}
	return c
}

func validateInferenceConfig(c *config.InferenceConfig) error {
	if c.APIVersion == "" {
		return errors.New("apiVersion is not defined")
	}

	if c.APIVersion != utils.APIv1alpha1 {
		return errors.Errorf("apiVersion %s is not supported", c.APIVersion)
	}

	if len(c.Models) == 0 {
		return errors.New("no models defined")
	}

	if len(c.Backends) > 1 {
		return errors.New("only one backend is supported at this time")
	}

	if slices.Contains(c.Backends, utils.BackendStableDiffusion) && (slices.Contains(c.Backends, utils.BackendExllama) || slices.Contains(c.Backends, utils.BackendExllamaV2)) {
		return errors.New("cannot specify both stablediffusion with exllama or exllama2 at this time")
	}

	if (slices.Contains(c.Backends, utils.BackendExllama) || slices.Contains(c.Backends, utils.BackendExllamaV2) || slices.Contains(c.Backends, utils.BackendMamba)) && c.Runtime != utils.RuntimeNVIDIA {
		return errors.New("exllama and mamba only supports nvidia cuda runtime. please add 'runtime: cuda' to your aikitfile.yaml")
	}

	backends := []string{utils.BackendExllama, utils.BackendExllamaV2, utils.BackendStableDiffusion, utils.BackendMamba}
	for _, b := range c.Backends {
		if !slices.Contains(backends, b) {
			return errors.Errorf("backend %s is not supported", b)
		}
	}

	runtimes := []string{"", utils.RuntimeNVIDIA, utils.RuntimeCPUAVX, utils.RuntimeCPUAVX2, utils.RuntimeCPUAVX512}
	if !slices.Contains(runtimes, c.Runtime) {
		return errors.Errorf("runtime %s is not supported", c.Runtime)
	}

	return nil
}

func DetectGitContext(ref string, keepGit bool) (llb.State, error) {
	g, err := gitutil.ParseGitRef(ref)
	if err != nil {
		return llb.State{}, err
	}
	commit := g.Commit
	if g.SubDir != "" {
		commit += ":" + g.SubDir
	}
	gitOpts := []llb.GitOption{dockerui.WithInternalName("load git source " + ref)}
	if keepGit {
		gitOpts = append(gitOpts, llb.KeepGitDir())
	}

	st := llb.Git(g.Remote, commit, gitOpts...)
	return st, nil
}

func DetectHTTPContext(ref string) (llb.State, string, bool) {
	filename := "context"
	httpPrefix := regexp.MustCompile(`^https?://`)
	if httpPrefix.MatchString(ref) {
		st := llb.HTTP(ref, llb.Filename(filename), dockerui.WithInternalName("load remote build context"))
		return st, filename, true
	}
	return llb.State{}, "", false
}

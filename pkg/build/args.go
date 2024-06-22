package build

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/sozercan/aikit/pkg/aikit/config"
	"github.com/sozercan/aikit/pkg/aikit2llb/inference"
	"github.com/sozercan/aikit/pkg/utils"
	"oras.land/oras-go/pkg/content"
	"oras.land/oras-go/pkg/oras"
	"oras.land/oras-go/pkg/registry/remote"
)

// parseBuildArgs parses the build arguments and configures inference settings.
func parseBuildArgs(opts map[string]string, inferenceCfg *config.InferenceConfig) error {
	if inferenceCfg == nil {
		return nil
	}

	// Get model and runtime arguments
	modelArg := getBuildArg(opts, "model")
	runtimeArg := getBuildArg(opts, "runtime")

	if modelArg != "" {
		var modelName, modelSource string
		var err error

		// Handle based on the URL prefix
		switch {
		case strings.HasPrefix(modelArg, "huggingface://"):
			// Handle Hugging Face URLs with optional branch
			modelSource, modelName, err = inference.ParseHuggingFaceURL(modelArg)
			if err != nil {
				return err
			}

		case strings.HasPrefix(modelArg, "http://"), strings.HasPrefix(modelArg, "https://"):
			// Handle HTTP(S) URLs directly
			modelName = utils.FileNameFromURL(modelArg)
			modelSource = modelArg

		case strings.HasPrefix(modelArg, "oci://"):
			// Handle OCI URLs
			modelName = parseOCIURL(modelArg)
			modelSource = modelArg
			if strings.Contains(modelArg, "registry.ollama.ai") {
				params := retrieveOllamaParams(modelArg)
			}

		default:
			// Assume it's a local file path
			modelName = path.Base(modelArg)
			modelSource = modelArg
		}

		// Set the inference configuration
		inferenceCfg.Runtime = runtimeArg
		inferenceCfg.Models = []config.Model{
			{
				Name:   modelName,
				Source: modelSource,
			},
		}
		inferenceCfg.Config = generateInferenceConfig(modelName)
	}

	return nil
}

// generateInferenceConfig generates the inference configuration for the given model name.
func generateInferenceConfig(modelName string) string {
	return fmt.Sprintf(`
- name: %[1]s
  backend: llama
  parameters:
    model: %[1]s
  stopwords:
%s`, modelName, )
}

// parseOCIURL extracts model name for OCI-based models.
func parseOCIURL(source string) string {
	const ollamaRegistryURL = "registry.ollama.ai"
	artifactURL := strings.TrimPrefix(source, "oci://")
	var modelName string

	if strings.HasPrefix(artifactURL, ollamaRegistryURL) {
		// Special handling for Ollama registry
		artifactURLWithoutTag := strings.Split(artifactURL, ":")[0]
		modelName = strings.Split(artifactURLWithoutTag, "/")[2]
	} else {
		// Generic OCI artifact
		modelName = path.Base(artifactURL)
		modelName = strings.Split(modelName, ":")[0]
		modelName = strings.Split(modelName, "@")[0]
	}
	return modelName
}

// formatStopwords formats the stopwords list for YAML output
func formatStopwords(stopwords []string) string {
	formatted := ""
	for _, word := range stopwords {
		formatted += fmt.Sprintf("    - %s\n", word)
	}
	return formatted
}

type OllamaConfig struct {
    NumCtx  int      `json:"num_ctx"`
    NumKeep int      `json:"num_keep"`
    Stop    []string `json:"stop"`
}

func retrieveOllamaParams(artifactURL string) OllamaConfig {
	artifactURLWithoutTag := strings.Split(artifactURL, ":")[0]
	tag := strings.Split(artifactURL, ":")[1]
	modelName := strings.Split(artifactURLWithoutTag, "/")[2]

	// Fetch the manifest and extract the digest
	digest, err := fetchOllamaDigest(modelName, tag)
	if err != nil {
		log.Fatalf("Error fetching manifest: %v", err)
	}

	err := fetchOllamaBlob("https://registry.ollama.ai", modelName, digest)


	//orasParamsCmd := fmt.Sprintf("oras blob fetch %[1]s@$(curl https://registry.ollama.ai/v2/library/%[2]s/manifests/%[3]s | jq -r '.layers[] | select(.mediaType == \"application/vnd.ollama.image.params\").digest') --output %[2]s", artifactURLWithoutTag, modelName, tag)

	// orasImage := "ghcr.io/oras-project/oras:v1.2.0"
	// toolingImage := llb.Image(orasImage) //, llb.Platform(platform)) //TODO

	// toolingImage = toolingImage.Run(utils.Shf("apk add jq curl && %s", orasParamsCmd)).Root()

}


func fetchOllamaBlob(registry string, repository string, digest string) error {
	// Construct the target registry URL
	url := fmt.Sprintf("%s/%s", registry, repository)

	// Create a new remote repository
	repo, err := remote.NewRepository(url)
	if err != nil {
		return fmt.Errorf("failed to create repository: %w", err)
	}

	// Configure authentication if needed (optional)
	// For this example, we are not adding authentication. You can configure
	// it using `repo.Client = auth.NewClient(...)` if authentication is required.

	// Set up an ORAS context
	ctx := context.Background()

	// Create a storage to save the blob
	memStore := content.NewMemory()

	// Fetch the blob
	_, err = oras.Copy(ctx, repo, digest, memStore, digest)
	if err != nil {
		return fmt.Errorf("failed to fetch blob: %w", err)
	}

	// Retrieve the blob content from the memory store
	blob, err := memStore.Fetch(ctx, digest)
	if err != nil {
		return fmt.Errorf("failed to retrieve blob: %w", err)
	}

	// Print the blob size and content (or handle it as needed)
	fmt.Printf("Blob size: %d bytes\n", len(blob))
	fmt.Printf("Blob content: %s\n", blob)

	return nil
}

// Layer represents the structure of a layer in the JSON response
type Layer struct {
	MediaType string `json:"mediaType"`
	Digest    string `json:"digest"`
}


// fetchOllamaDigest fetches the manifest from the provided URL and extracts the required digest
func fetchOllamaDigest(library string, model string) (string, error) {
	// Construct the URL
	url := fmt.Sprintf("https://registry.ollama.ai/v2/library/%s/manifests/%s", library, model)

	// Perform the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch manifest: %s", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the JSON response
	var result struct {
		Layers []Layer `json:"layers"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	// Find the layer with the specified media type
	for _, layer := range result.Layers {
		if layer.MediaType == "application/vnd.ollama.image.params" {
			return layer.Digest, nil
		}
	}

	return "", fmt.Errorf("no layer found with the specified media type")
}
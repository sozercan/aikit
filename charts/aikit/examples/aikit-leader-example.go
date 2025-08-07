package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"k8s.io/klog/v2"
)

// LocalAIConfig represents the configuration structure for LocalAI
type LocalAIConfig struct {
	Debug     bool          `yaml:"debug"`
	Address   string        `yaml:"address"`
	Port      int          `yaml:"port"`
	Models    []ModelConfig `yaml:"models"`
	Backends  []string      `yaml:"backends,omitempty"`
	Workers   []WorkerConfig `yaml:"workers,omitempty"`
}

type ModelConfig struct {
	Name       string            `yaml:"name"`
	Model      string            `yaml:"model"`
	Backend    string            `yaml:"backend,omitempty"`
	Parameters map[string]interface{} `yaml:"parameters,omitempty"`
}

type WorkerConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	klog.Info("Starting AIKit Leader for distributed inference")

	// Get LeaderWorkerSet configuration from environment
	lwsSize := getLWSGroupSize()
	lwsLeaderAddress := os.Getenv("LWS_LEADER_ADDRESS")

	klog.Infof("LWS Group Size: %d, Leader Address: %s", lwsSize, lwsLeaderAddress)

	// If single node, run LocalAI directly
	if lwsSize <= 1 {
		klog.Info("Single node mode, starting LocalAI directly")
		return startLocalAI(ctx, nil)
	}

	// Discover worker pods
	klog.Info("Discovering worker pods...")
	workers, err := discoverWorkers(lwsSize, lwsLeaderAddress)
	if err != nil {
		return fmt.Errorf("failed to discover workers: %w", err)
	}

	klog.Infof("Discovered %d workers: %v", len(workers), workers)

	// Generate distributed LocalAI configuration
	if err := generateDistributedConfig(workers); err != nil {
		return fmt.Errorf("failed to generate distributed config: %w", err)
	}

	// Start LocalAI with distributed configuration
	klog.Info("Starting LocalAI with distributed configuration")
	return startLocalAI(ctx, workers)
}

func getLWSGroupSize() int {
	if s := os.Getenv("LWS_GROUP_SIZE"); s != "" {
		if v, err := strconv.Atoi(s); err == nil {
			return v
		}
	}
	return 1 // fallback to single node
}

func discoverWorkers(lwsSize int, lwsLeaderAddress string) ([]WorkerConfig, error) {
	if lwsLeaderAddress == "" {
		return nil, fmt.Errorf("LWS_LEADER_ADDRESS not set")
	}

	serviceTokens := strings.Split(lwsLeaderAddress, ".")
	if len(serviceTokens) < 2 {
		return nil, fmt.Errorf("invalid LWS_LEADER_ADDRESS format: %s", lwsLeaderAddress)
	}

	var workers []WorkerConfig

	// Workers are indexed 1, 2, 3... (0 is leader)
	for i := 1; i < lwsSize; i++ {
		workerHost := fmt.Sprintf("%s-%d.%s",
			serviceTokens[0], i, strings.Join(serviceTokens[1:], "."))

		klog.Infof("Waiting for worker %d at %s", i, workerHost)

		// Wait for worker to be resolvable
		if ip := waitForWorker(workerHost); ip != "" {
			workers = append(workers, WorkerConfig{
				Address: ip,
				Port:    50051, // Default worker port
			})
			klog.Infof("Worker %d ready at %s", i, ip)
		} else {
			klog.Warningf("Worker %d at %s not ready, continuing without it", i, workerHost)
		}
	}

	return workers, nil
}

func waitForWorker(host string) string {
	maxAttempts := 30
	for attempt := 0; attempt < maxAttempts; attempt++ {
		if ips, err := net.LookupIP(host); err == nil && len(ips) > 0 {
			return ips[0].String()
		}
		klog.V(2).Infof("Attempt %d/%d: waiting for %s", attempt+1, maxAttempts, host)
		time.Sleep(5 * time.Second)
	}
	return ""
}

func generateDistributedConfig(workers []WorkerConfig) error {
	// Read existing config if it exists
	var config LocalAIConfig
	if configFile := os.Getenv("CONFIG_FILE_PATH"); configFile != "" {
		// TODO: Load existing configuration
		klog.Infof("Loading base config from %s", configFile)
	}

	// Set distributed defaults
	config.Debug = true
	config.Address = "0.0.0.0"
	config.Port = 8080
	config.Workers = workers

	// Add model configuration
	if modelPath := os.Getenv("MODEL_PATH"); modelPath != "" {
		config.Models = []ModelConfig{
			{
				Name:    "distributed-model",
				Model:   modelPath,
				Backend: "distributed", // Hypothetical distributed backend
				Parameters: map[string]interface{}{
					"workers": len(workers),
					"distributed": true,
				},
			},
		}
	}

	// Write configuration to file
	configData, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	configPath := "/tmp/distributed-config.json"
	if err := os.WriteFile(configPath, configData, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	klog.Infof("Generated distributed config at %s", configPath)
	return nil
}

func startLocalAI(ctx context.Context, workers []WorkerConfig) error {
	args := []string{
		"--address=0.0.0.0",
		"--port=8080",
	}

	// Add distributed configuration if workers are present
	if len(workers) > 0 {
		args = append(args, "--config-file=/tmp/distributed-config.json")

		// NOTE: These are hypothetical flags - LocalAI may not support these
		args = append(args, "--distributed=true")
		args = append(args, fmt.Sprintf("--workers=%d", len(workers)))
	}

	// Add debug flag if requested
	if os.Getenv("DEBUG") == "true" {
		args = append(args, "--debug")
	}

	klog.Infof("Starting LocalAI with args: %v", args)

	cmd := exec.CommandContext(ctx, "/usr/bin/local-ai", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start LocalAI: %w", err)
	}

	// Wait for the process to complete
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("LocalAI exited with error: %w", err)
	}

	return nil
}

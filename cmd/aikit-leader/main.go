package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	yaml "gopkg.in/yaml.v2"
	"k8s.io/klog/v2"
)

// LocalAIConfig represents a minimal configuration structure for LocalAI.
// This assumes LocalAI can take a YAML config with address/port and models.
// Adjust as needed to match actual LocalAI config schema.
type LocalAIConfig struct {
	Debug   bool          `yaml:"debug,omitempty"`
	Address string        `yaml:"address,omitempty"`
	Port    int           `yaml:"port,omitempty"`
	Models  []ModelConfig `yaml:"models,omitempty"`
	// Optional: a hypothetical distributed section. Keep loose to avoid coupling.
	Distributed *DistributedConfig `yaml:"distributed,omitempty"`
}

type ModelConfig struct {
	Name       string                 `yaml:"name"`
	Model      string                 `yaml:"model"`
	Backend    string                 `yaml:"backend,omitempty"`
	Parameters map[string]interface{} `yaml:"parameters,omitempty"`
}

type DistributedConfig struct {
	Enabled bool           `yaml:"enabled"`
	Workers []WorkerConfig `yaml:"workers"`
}

type WorkerConfig struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

func main() {
	klog.InitFlags(nil)
	defer klog.Flush()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := run(ctx); err != nil {
		klog.ErrorS(err, "aikit-leader failed")
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	klog.Info("Starting aikit-leader for distributed inference")

	// Read environment
	lwsSize := getenvInt("LWS_GROUP_SIZE", 1)
	lwsLeaderAddress := os.Getenv("LWS_LEADER_ADDRESS")
	leaderPort := getenvInt("LEADER_PORT", 8080)
	workerPort := getenvInt("WORKER_PORT", 50051)
	modelPath := os.Getenv("MODEL_PATH")
	debug := os.Getenv("DEBUG") == "true"

	klog.Infof("LWS_GROUP_SIZE=%d, LWS_LEADER_ADDRESS=%q, leaderPort=%d, workerPort=%d", lwsSize, lwsLeaderAddress, leaderPort, workerPort)

	// Discover workers if this is a multi-node setup
	var workers []WorkerConfig
	if lwsSize > 1 && lwsLeaderAddress != "" {
		// Use a timeout context for worker discovery to avoid blocking indefinitely
		discoveryCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
		defer cancel()

		var err error
		workers, err = discoverWorkers(discoveryCtx, lwsSize, lwsLeaderAddress, workerPort)
		if err != nil {
			klog.Warningf("Worker discovery failed (will start LocalAI anyway): %v", err)
		} else {
			klog.Infof("Discovered %d workers", len(workers))
		}
	} else {
		klog.Info("Single node mode or no leader address - starting LocalAI directly")
	}

	// Build LocalAI args - use 'run' command for LocalAI v2.26+
	args := []string{
		"run",
		fmt.Sprintf("--address=:%d", leaderPort),
	}
	if debug {
		args = append(args, "--log-level=debug")
	}

	// Add model path if specified
	if modelPath != "" {
		args = append(args, "--models-path="+modelPath)
	}

	// Start LocalAI as child process
	cmdPath := getenvString("LOCALAI_BINARY", "/usr/bin/local-ai")
	klog.Infof("Starting %s %v", cmdPath, args)

	cmd := exec.CommandContext(ctx, cmdPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start local-ai: %w", err)
	}

	// Wait until process exits or context canceled
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	select {
	case <-ctx.Done():
		// Try to terminate gracefully
		_ = cmd.Process.Signal(syscall.SIGTERM)
		select {
		case err := <-done:
			return err
		case <-time.After(10 * time.Second):
			_ = cmd.Process.Kill()
			return ctx.Err()
		}
	case err := <-done:
		return err
	}
}

func discoverWorkers(ctx context.Context, lwsSize int, lwsLeaderAddress string, port int) ([]WorkerConfig, error) {
	if lwsLeaderAddress == "" {
		return nil, fmt.Errorf("LWS_LEADER_ADDRESS not set")
	}
	parts := strings.Split(lwsLeaderAddress, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid LWS_LEADER_ADDRESS: %q", lwsLeaderAddress)
	}

	workers := make([]WorkerConfig, 0, max(0, lwsSize-1))
	for i := 1; i < lwsSize; i++ { // 0 is leader
		host := fmt.Sprintf("%s-%d.%s", parts[0], i, strings.Join(parts[1:], "."))
		klog.Infof("Looking for worker %d at %s", i, host)
		ip := waitForDNS(ctx, host, 12, 5*time.Second) // Reduced from 30 attempts
		if ip == "" {
			klog.Warningf("worker %d not reachable at %s", i, host)
			continue
		}
		workers = append(workers, WorkerConfig{Address: ip, Port: port})
		klog.Infof("worker %d ready at %s:%d", i, ip, port)
	}

	// Return success even if no workers found - this allows leader to start anyway
	if len(workers) == 0 {
		klog.Info("No workers discovered, will run in single-node mode")
	}
	return workers, nil
}

func waitForDNS(ctx context.Context, host string, attempts int, delay time.Duration) string {
	for a := 0; a < attempts; a++ {
		ips, err := net.LookupIP(host)
		if err == nil && len(ips) > 0 {
			return ips[0].String()
		}
		select {
		case <-ctx.Done():
			return ""
		case <-time.After(delay):
		}
	}
	return ""
}

func writeYAML(path string, v interface{}) error {
	data, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func getenvInt(key string, def int) int {
	s := os.Getenv(key)
	if s == "" {
		return def
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return v
}

func getenvString(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

# AIKit Development Instructions

Always reference these instructions first and fallback to search or bash commands only when you encounter unexpected information that does not match the info here.

## Working Effectively

AIKit is a Go-based platform for building and deploying AI model containers. It uses Docker for containerization and supports CPU, GPU (NVIDIA CUDA), and Apple Silicon platforms.

### Prerequisites and Setup
- Install Go 1.24.4 or later: `go version` should show 1.24+
- Install Docker: `docker --version` 
- Install golangci-lint v2.1.6: `curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.1.6`
- Verify Go dependencies: `go mod download && go mod verify`

### Development Workflow
- Run unit tests: `make test` -- takes 3 seconds. NEVER CANCEL.
- Run linting: `export PATH="$(go env GOPATH)/bin:$PATH" && golangci-lint run -v ./... --timeout 5m` -- takes 1 second. NEVER CANCEL.
- Check code formatting: golangci-lint automatically handles gofmt, gofumpt, goimports
- View test coverage: Tests generate `coverage.txt` with 90% coverage on config, 21% on build packages

### Build Process **IMPORTANT LIMITATIONS**
- `make build-aikit` -- **FAILS in sandboxed environments due to TLS certificate issues**
- In normal environments: expected to take 45+ minutes based on CI timeouts. NEVER CANCEL. Set timeout to 60+ minutes.
- `make build-test-model` -- **REQUIRES successful aikit build first**
- In normal environments: expected to take 60+ minutes for model builds. NEVER CANCEL. Set timeout to 90+ minutes.

### Test Model Development **LIMITATION**
- Local model testing: `make run-test-model` -- **CANNOT BE VALIDATED** due to build limitations
- GPU testing: `make run-test-model-gpu` -- **CANNOT BE VALIDATED** due to build limitations
- Apple Silicon: `make run-test-model-applesilicon` -- **CANNOT BE VALIDATED** due to build limitations

### Validation **CRITICAL**
**NEVER CANCEL ANY BUILD OR TEST COMMAND** - Builds may take 45+ minutes, tests may take 15+ minutes in CI environments.
**MANUAL VALIDATION REQUIREMENT**: After building and running applications in normal environments, you MUST test:
1. Navigate to `http://localhost:8080/chat` to verify WebUI
2. Test OpenAI-compatible API:
```bash
curl http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d '{
  "model": "llama-3.2-1b-instruct", 
  "messages": [{"role": "user", "content": "explain kubernetes in a sentence"}]
}'
```
3. Verify response contains `choices` array with assistant message

### Pre-commit Quality Checks
- Always run before committing: `make test && golangci-lint run -v ./... --timeout 5m`
- CI requires all linting checks to pass (lint.yaml workflow)
- Pre-commit hooks available: `pre-commit install` (requires pip install pre-commit)

## Repository Structure

### Key Directories
- `cmd/frontend/`: Main AIKit CLI application entry point
- `pkg/aikit/config/`: Configuration parsing for aikitfiles (90% test coverage)
- `pkg/aikit2llb/`: BuildKit LLB conversion for Docker builds
- `pkg/build/`: Build logic and validation (21% test coverage) 
- `pkg/utils/`: Utility functions (42% test coverage)
- `test/`: Test aikitfile configurations (aikitfile-*.yaml)
- `models/`: Pre-made model configurations
- `charts/`: Kubernetes Helm charts for deployment
- `website/`: Documentation website (Docusaurus)

### Important Files
- `Makefile`: All build targets and commands
- `.golangci.yaml`: Linting configuration (requires golangci-lint v2)
- `go.mod`: Go 1.23.0 with toolchain go1.24.4
- `Dockerfile`: AIKit binary build (fails in sandboxed environments)
- `test/aikitfile-llama.yaml`: Default test model configuration

### Build Configuration Files
- `aikitfile-*.yaml`: Model-specific build configurations in test/
- `Dockerfile.base`: Base image for model containers
- `Dockerfile.base-applesilicon`: Apple Silicon specific base

## Common Workflows

### Adding New Model Support
1. Create new `aikitfile-[model].yaml` in test/ directory
2. Define model source, SHA256, and prompt templates
3. Configure backend (llama, exllama2, diffusers, etc.)
4. Test build: `make build-test-model TEST_FILE=test/aikitfile-[model].yaml`
5. Test locally: `make run-test-model`

### Development Testing
1. Run unit tests: `make test` (3 seconds)
2. Run linting: `golangci-lint run -v ./... --timeout 5m` (1 second)
3. **Build validation cannot be completed in sandboxed environments**

### CI/CD Integration
- GitHub Actions workflows expect 240-minute timeouts for Docker builds
- Unit tests run in ~1 minute in CI
- Multi-platform builds (linux/amd64, linux/arm64) supported
- GPU testing requires NVIDIA Docker runtime

## Troubleshooting

### Known Issues
- **TLS Certificate Issues**: Docker builds fail in sandboxed environments with "certificate signed by unknown authority" errors
- **Workaround**: This is expected in restricted environments; builds work in normal development setups
- **golangci-lint Version**: Must use v2.1.6 specifically for v2 configuration file

### Build Failures
- If `make build-aikit` fails with TLS errors: Expected in sandboxed environments
- If linting fails: Check golangci-lint version with `golangci-lint version`
- If tests fail: Run `go mod tidy` first, then `make test`

### Performance Expectations
- Unit tests: 3 seconds
- Linting: 1 second  
- AIKit binary build: 45-60 minutes (in normal environments)
- Model container builds: 60-90 minutes (in normal environments)
- **NEVER CANCEL** long-running builds - they complete successfully given sufficient time

## Platform-Specific Notes

### CPU Support
- Supports AMD64 and ARM64 architectures
- Automatic instruction set optimization
- GGUF model format recommended

### GPU Support  
- NVIDIA CUDA runtime required: `--gpus all` flag
- GPTQ and EXL2 models via exllama2 backend
- Test with: `make run-test-model-gpu`

### Apple Silicon
- Experimental support via Podman
- Only GGUF models supported
- Test with: `podman run --device /dev/dri`

**Note**: All container execution validation cannot be completed in current environment due to build limitations.
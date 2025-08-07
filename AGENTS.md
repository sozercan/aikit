# AIKit Developer Guide

This document gives you the fast inner dev loop (tests + lint) and the slower container/model build loop. Start here before ad‑hoc exploration.

## 1. Quick Start (Fast Loop)

```bash
make test          # run unit tests (-race + coverage)
make lint          # golangci-lint (format + static checks)
make build-aikit   # optional: build aikit binary container (slow)
make build-test-model TEST_FILE=test/aikitfile-llama.yaml  # build sample model image (slow)
make run-test-model # run the model API locally (after successful build)
```

Visit: <http://localhost:8080/chat>

Validate API (OpenAI compatible):

```bash
curl -s http://localhost:8080/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{"model":"llama-3.2-1b-instruct","messages":[{"role":"user","content":"explain kubernetes in a sentence"}]}' | jq '.choices[0].message.content'
```

## 2. Prerequisites

- Go: go.mod specifies `go 1.23.0` with toolchain `go1.24.4` (install Go >=1.23; `go version` should reflect the toolchain).
- Docker with BuildKit / buildx: `docker buildx ls` must show a builder.
- (Optional) NVIDIA GPU runtime for GPU model execution.
- golangci-lint v2.x (CI uses 2.1.6 or later):

  ```bash
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
    sh -s -- -b "$(go env GOPATH)/bin" v2.1.6
  ```

## 3. Concepts

- aikitfile: YAML describing model source + backend/runtime parameters.
- Fast loop: Go code changes → `make test` + `make lint` (seconds to <1 min locally).
- Slow loop: Building containers / models (tens of minutes depending on model + platform + network cache).
- Backends: llama, exllama2 (GPTQ/EXL2), diffusers (image), mamba, etc.

## 4. Make Targets (Selected)

| Target                                    | Purpose                                      | Notes                                           |
| ----------------------------------------- | -------------------------------------------- | ----------------------------------------------- |
| `make test`                               | Run all tests with race + coverage           | Produces coverage.txt                           |
| `make lint`                               | Run golangci-lint (format + static analysis) | Uses `.golangci.yaml`                           |
| `make build-aikit`                        | Build base aikit image                       | Slow; may fail behind restrictive proxies (TLS) |
| `make build-test-model`                   | Build model image from `TEST_FILE`           | Set `TEST_FILE`, `TEST_IMAGE_NAME`, `RUNTIME`   |
| `make run-test-model`                     | Run model container (CPU default)            | Exposes :8080                                   |
| `make run-test-model-gpu`                 | Run with GPU                                 | Requires `--gpus all` support                   |
| `make run-test-model-applesilicon`        | Run via Podman on Apple Silicon              | Experimental                                    |
| `make release-manifest NEWVERSION=vX.Y.Z` | Bump version & chart                         | Updates Makefile + Chart                        |

Environment overrides:

```text
REGISTRY (default ghcr.io/kaito-project)
TAG (default test)
TEST_FILE (default test/aikitfile-llama.yaml)
TEST_IMAGE_NAME (default testmodel)
OUTPUT_TYPE (e.g. type=docker or type=registry)
PLATFORMS (default linux/amd64,linux/arm64)
RUNTIME (backend-specific, passed as build arg)
```

## 5. Adding a New Model

There are two paths depending on intent:

1. Local / experimental test model (not published)
2. Pre-made (shared) model configuration contributed under `models/`

### 5.1 Test Model (CI / Validation)

Files under `test/` are canonical test model definitions exercised in CI to validate parsing, build translation, and runtime behaviors across backends. They are not "temporary" or destined for promotion; they should remain stable and minimal.

| Aspect             | Local Test Model                                                  |
| ------------------ | ----------------------------------------------------------------- |
| File location      | `test/` directory                                                 |
| Filename pattern   | `aikitfile-<shortname>.yaml`                                      |
| Syntax directive   | `#syntax=aikit:test` (lightweight base)                           |
| Build command      | `make build-test-model TEST_FILE=test/aikitfile-<shortname>.yaml` |
| Image tag override | `TAG=<tag>` (optional)                                            |
| Commit?            | Optional (only if useful for others)                              |

Steps (adding or adjusting a CI test model):

1. Start from a similar small model: `cp test/aikitfile-llama.yaml test/aikitfile-<shortname>.yaml`.
2. Keep artifact size modest (fast download) and pick a stable upstream URL.
3. Update `source`, `sha256`, prompt templates, and backend config.
4. Build locally (if environment allows) using: `make build-test-model TEST_FILE=test/aikitfile-<shortname>.yaml`.
5. Run locally (optional in restrictive envs): `make run-test-model TEST_FILE=...` (set TAG/RUNTIME if needed).
6. Ensure changes do not introduce excessive build time or large layers.
7. Commit the new/changed test file; adjust or add unit tests if specific behavior needs validation.

### 5.2 Pre‑Made Shared Model (Distributed)

Use this when adding a model that should ship with the project (discoverable & maintained).

| Aspect           | Pre-made Shared Model                                   |
| ---------------- | ------------------------------------------------------- |
| File location    | `models/` directory                                     |
| Filename         | `<model-name>.yaml` (e.g. `llama-3.2-1b-instruct.yaml`) |
| Syntax directive | `#syntax=ghcr.io/kaito-project/aikit/aikit:latest`      |
| Naming           | Match canonical upstream model name (size + type)       |
| Commit?          | Yes (PR required)                                       |
| Docs update      | Add / adjust `website/docs/premade-models.md` if needed |

Steps:

1. Start from a similar model in `models/` (choose same backend family).
2. Keep `#syntax=ghcr.io/kaito-project/aikit/aikit:latest` unless a version pin is justified.
3. Add / verify `sha256` for the artifact (GGUF / safetensors / etc.).
4. Ensure prompt templates include at least: `chatMsg` (or equivalent), and any function / completion variants required by the backend.
5. Verify backend block in `config:` aligns (e.g. `backend: llama`, parameters model filename matches downloaded file name).
6. (Optional) Build locally first using: `make build-test-model TEST_FILE=models/<model-name>.yaml TEST_IMAGE_NAME=<shorttag>`.
7. Run & validate (same as test model path).
8. Add entry to docs if it's a new family / size not already listed.
9. Open PR with rationale (upstream source, license compatibility, checksum).

### 5.3 Key Differences Summary

| Dimension             | Test Model (`test/`)                                      | Pre-made (`models/`)                         |
| --------------------- | --------------------------------------------------------- | -------------------------------------------- |
| Purpose               | Rapid local experiment                                    | Distributed curated config                   |
| Naming                | `aikitfile-*.yaml`                                        | `<model>.yaml`                               |
| Base syntax image     | `aikit:test`                                              | Published `aikit:latest`                     |
| Required to commit    | No                                                        | Yes                                          |
| Documentation change  | No                                                        | Maybe (if new)                               |
| Build command example | `make build-test-model TEST_FILE=test/aikitfile-foo.yaml` | Same command but `TEST_FILE=models/foo.yaml` |

### 5.4 Choosing Where to Add a Model

Place in `test/` when:

- Primary purpose is CI coverage (exercise backend/path/runtime variation).
- Artifact is small enough to keep build times reasonable.
- Frequent updates are expected while iterating on support.

Place in `models/` when:

- Intended for end-user consumption and documentation.
- Represents a widely used or reference model.
- You will maintain checksum and prompt template quality over time.

## 6. Running (Platforms)

CPU (AMD64/ARM64): `make run-test-model`

GPU (NVIDIA): `make run-test-model-gpu` (requires Docker runtime w/ CUDA, driver + nvidia-container-toolkit installed).

Apple Silicon (experimental via Podman): `make run-test-model-applesilicon` (GGUF models only, uses `--device /dev/dri`).

## 7. Repository Map

- `cmd/frontend/` – main entrypoint.
- `pkg/aikit/config/` – aikitfile parsing & validation.
- `pkg/aikit2llb/` – BuildKit LLB conversion logic.
- `pkg/build/` – build orchestration & validation.
- `pkg/utils/` – shared helpers.
- `pkg/version/` – version constant.
- `models/` – pre-made model configs (GGUF/GPTQ/EXL2/diffusers).
- `test/` – sample aikitfiles used in builds/tests.
- `charts/` – Helm chart for deployment.
- `website/` – docs site (Docusaurus).

Key files:

- `Makefile`, `.golangci.yaml`, `go.mod`, `Dockerfile`, `Dockerfile.base*`.

## 8. Validation Checklist

After a successful model image build & run:

1. Web UI: <http://localhost:8080/chat> loads.
2. `/v1/chat/completions` returns JSON with `choices[0].message`.
3. (Optional) List models endpoint (if implemented) returns your model ID.

## 9. CI Overview

- Lint + tests run on every PR (golangci-lint v2.x + `go test`).
- Docker/model builds use generous timeouts (multi-platform buildx).
- Helm chart & version updated via `make release-manifest` in release workflow.

## 10. Troubleshooting

TLS / certificate errors during `docker buildx build`:

- Cause: restricted / sandbox environment trust store mismatch.
- Action: retry on a standard dev host or configure corporate CA into Docker.

Lint failures:

- Confirm version: `golangci-lint version` (ensure v2.x).
- Run individually to isolate: `golangci-lint run --disable-all -E <linter> ./...`.

Test failures after dependency changes:

- `go mod tidy && make test`.

Slow builds / timeouts:

- Ensure BuildKit: `export DOCKER_BUILDKIT=1` (if older Docker setups).
- Leverage layer cache by not cleaning local Docker state unnecessarily.

GPU not detected:

- Check: `docker run --gpus all --rm nvidia/cuda:12.4.0-base nvidia-smi`.

Apple Silicon issues:

- Verify Podman version & `--device /dev/dri` availability.

## 11. Style & Quality

- Always run `make test && make lint` before pushing.
- Formatting enforced by golangci-lint (gofmt/gofumpt/goimports/gci).
- Coverage report: `coverage.txt` (opened by many IDEs); thresholds not enforced—focus on critical packages first.

## 12. Releasing

1. Update version: `make release-manifest NEWVERSION=vX.Y.Z`.
2. Commit & tag: `git commit -am "release: vX.Y.Z" && git tag vX.Y.Z`.
3. Push: `git push origin main --tags`.
4. CI publishes updated chart & images (workflow dependent).

## 13. Contributing

- Fork / branch, follow conventional concise commit messages.
- Keep PRs focused (one feature/fix).
- Include tests where logic changes.
- Update docs (`website/docs/*.md`) if behavior or flags change.

## 14. Sandbox Limitations

If you are reading this inside a restricted or ephemeral environment:

- Large Docker builds may fail (TLS) or be too slow; defer to a full dev machine.
- API runtime validation depends on successful local image builds.

## 15. FAQ (Quick)

Q: Can I skip `make build-aikit` before a model build?
A: Yes; they are independent. `build-test-model` uses the aikitfile directly.

Q: How do I change model backend?
A: Set the appropriate backend fields in the aikitfile (see examples in `test/`).

Q: Where do runtime-specific optimizations happen?
A: During the Docker build via BuildKit args (`runtime`, platform target, model format).

---
Keep this guide concise—avoid adding transient metrics (exact minutes, coverage %) unless enforced by CI.

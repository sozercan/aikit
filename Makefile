REGISTRY ?= ghcr.io/sozercan
KIND_VERSION ?= 0.20.0
KUBERNETES_VERSION ?= 1.29.0
TAG ?= test
OUTPUT_TYPE ?= type=docker
TEST_IMAGE_NAME ?= testmodel
TEST_FILE ?= test/aikitfile-llama.yaml

GIT_COMMIT := $(shell git rev-list --abbrev-commit --tags --max-count=1)
GIT_TAG := $(shell git describe --abbrev=0 --tags ${GIT_COMMIT} 2>/dev/null || true)
LDFLAGS := "-X github.com/sozercan/aikit/pkg/version.Version=$(GIT_TAG:%=%)"

.PHONY: lint
lint:
	golangci-lint run -v ./... --timeout 5m

.PHONY: build-aikit
build-aikit:
	docker buildx build . -t ${REGISTRY}/aikit:${TAG} --output=${OUTPUT_TYPE} --build-arg LDFLAGS=${LDFLAGS}

.PHONY: build-test-model
build-test-model:
	docker buildx build . -t ${REGISTRY}/${TEST_IMAGE_NAME}:${TAG} -f ${TEST_FILE} --output=${OUTPUT_TYPE} \
		--progress=plain --provenance=false

.PHONY: run-test-model
run-test-model:
	docker run --rm -p 8080:8080 ${REGISTRY}/${TEST_IMAGE_NAME}:${TAG}

.PHONY: run-test-model-gpu
run-test-model-gpu:
	docker run --rm -p 8080:8080 --gpus all ${REGISTRY}/${TEST_IMAGE_NAME}:${TAG}

.PHONY: test
test:
	go test -v ./... -race -coverprofile=coverage.txt -covermode=atomic

.PHONY: test-e2e-dependencies
test-e2e-dependencies:
	mkdir -p ${GITHUB_WORKSPACE}/bin
	echo "${GITHUB_WORKSPACE}/bin" >> ${GITHUB_PATH}

	# used for kubernetes test
	curl -sSL https://dl.k8s.io/release/v${KUBERNETES_VERSION}/bin/linux/amd64/kubectl -o ${GITHUB_WORKSPACE}/bin/kubectl && chmod +x ${GITHUB_WORKSPACE}/bin/kubectl
	curl -sSL https://github.com/kubernetes-sigs/kind/releases/download/v${KIND_VERSION}/kind-linux-amd64 -o ${GITHUB_WORKSPACE}/bin/kind && chmod +x ${GITHUB_WORKSPACE}/bin/kind

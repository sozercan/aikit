REGISTRY ?= ghcr.io/sozercan
BATS_TESTS_FILE ?= test/bats/test.bats
BATS_VERSION ?= 1.10.0
KIND_VERSION ?= 0.20.0
KUBERNETES_VERSION ?= 1.28.0
TAG ?= test
OUTPUT_TYPE ?= type=docker
TEST_FILE ?= test/aikitfile-llama.yaml
PULL ?=
NO_CACHE ?=

.PHONY: lint
lint:
	golangci-lint run -v ./... --timeout 5m

.PHONY: build-aikit
build-aikit:
	docker buildx build . -t ${REGISTRY}/aikit:${TAG} --output=${OUTPUT_TYPE}

.PHONY: build-test-model
build-test-model:
	docker buildx build . -t ${REGISTRY}/testmodel:${TAG} -f ${TEST_FILE} --output=${OUTPUT_TYPE} --progress plain

.PHONY: run-test-model
run-test-model:
	docker run -p 8080:8080 ${REGISTRY}/testmodel:${TAG}

.PHONY: test
test:
	go test -v ./... -race -coverprofile=coverage.txt -covermode=atomic

.PHONY: test-e2e-dependencies
test-e2e-dependencies:
	mkdir -p ${GITHUB_WORKSPACE}/bin
	echo "${GITHUB_WORKSPACE}/bin" >> ${GITHUB_PATH}

	curl -sSLO https://github.com/bats-core/bats-core/archive/v${BATS_VERSION}.tar.gz && tar -zxvf v${BATS_VERSION}.tar.gz && bash bats-core-${BATS_VERSION}/install.sh ${GITHUB_WORKSPACE}

	# used for kubernetes test
	curl -sSL https://dl.k8s.io/release/v${KUBERNETES_VERSION}/bin/linux/amd64/kubectl -o ${GITHUB_WORKSPACE}/bin/kubectl && chmod +x ${GITHUB_WORKSPACE}/bin/kubectl
	curl -sSL https://github.com/kubernetes-sigs/kind/releases/download/v${KIND_VERSION}/kind-linux-amd64 -o ${GITHUB_WORKSPACE}/bin/kind && chmod +x ${GITHUB_WORKSPACE}/bin/kind

.PHONY: test-e2e
test-e2e:
	/home/runner/work/aikit/aikit/bin/bats --verbose-run --trace ${BATS_TESTS_FILE}

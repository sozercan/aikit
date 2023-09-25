REGISTRY ?= docker.io/sozercan
BATS_TESTS_FILE ?= test/bats/test.bats
BATS_VERSION ?= 1.10.0
KIND_VERSION ?= 0.20.0
KUBERNETES_VERSION ?= 1.28.0

.PHONY: lint
lint:
	golangci-lint run -v ./... --timeout 5m

.PHONY: build-aikit
build-aikit:
	docker buildx build . -t ${REGISTRY}/aikit:latest --push --pull --no-cache

.PHONY: build-test-model
build-test-model:
	docker buildx build . -t ${REGISTRY}/testmodel:latest -f test/aikitfile.yaml --push --pull --no-cache

.PHONY: run-test-model
run-test-model:
	docker run --pull always -p 8080:8080 ${REGISTRY}/testmodel:latest

.PHONY: test
test:
	go test -v ./... -race -coverprofile=coverage.txt -covermode=atomic

.PHONY: test-e2e-dependencies
test-e2e-dependencies:
	mkdir -p ${GITHUB_WORKSPACE}/bin
	echo "${GITHUB_WORKSPACE}/bin" >> ${GITHUB_PATH}

	curl -sSL https://dl.k8s.io/release/v${KUBERNETES_VERSION}/bin/linux/amd64/kubectl -o ${GITHUB_WORKSPACE}/bin/kubectl && chmod +x ${GITHUB_WORKSPACE}/bin/kubectl

	curl -sSL https://github.com/kubernetes-sigs/kind/releases/download/v${KIND_VERSION}/kind-linux-amd64 -o ${GITHUB_WORKSPACE}/bin/kind && chmod +x ${GITHUB_WORKSPACE}/bin/kind

	curl -sSLO https://github.com/bats-core/bats-core/archive/v${BATS_VERSION}.tar.gz && tar -zxvf v${BATS_VERSION}.tar.gz && bash bats-core-${BATS_VERSION}/install.sh ${GITHUB_WORKSPACE}

.PHONY: test-e2e
test-e2e:
	/home/runner/work/aikit/aikit/bin/bats --verbose-run --trace ${BATS_TESTS_FILE}

# build-aikit:
# 	docker buildx build . -t sozercan/aikit:latest --push

# build-llm:
# 	docker buildx build . -t sozercan/myllm:latest -f test/aikitfile.yaml --push

# run:
# 	docker run -p 8080:8080 sozercan/myllm:latest

# test:
# 	go test -v ./... -race -coverprofile=coverage.txt -covermode=atomic

BATS_VERSION ?= 1.10.0
BATS_TESTS_FILE ?= test/bats/test.bats

.PHONY: lint
lint:
	golangci-lint run -v ./... --timeout 5m

.PHONY: build-aikit
build-aikit:
	docker buildx build . -t sozercan/aikit:latest --push --pull --no-cache

.PHONY: build-test-model
build-test-model:
	docker buildx build . -t sozercan/testmodel:latest -f test/aikitfile.yaml --push --pull --no-cache

.PHONY: run-test-model
run-test-model:
	docker run -p 8080:8080 --pull always sozercan/testmodel:latest

.PHONY: test
test:
	go test -v ./... -race -coverprofile=coverage.txt -covermode=atomic

.PHONY: test-e2e-dependencies
test-e2e-dependencies:
	curl -sSLO https://github.com/bats-core/bats-core/archive/v${BATS_VERSION}.tar.gz && tar -zxvf v${BATS_VERSION}.tar.gz && bash bats-core-${BATS_VERSION}/install.sh ${GITHUB_WORKSPACE}

.PHONY: test-e2e
test-e2e:
	bats -t ${BATS_TESTS_FILE}

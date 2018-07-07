GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)


.PHONY: all build ci clean dep ginkgo test install

CERT_FOLDER = data/certs
GO_PROJECT = github.com/sks/mqttfaas
BUILD_DEST = build
COMMIT_HASH=`git rev-parse --short HEAD`
GOFILES_NOVENDOR=`find . -type f -name '*.go' -not -path "./vendor/*"`
GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_DIRTY  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

LDFLAGS += -w -s -extldflags -static

DOCKER_BUILD_ARGS=--build-arg http_proxy=${HTTP_PROXY} --build-arg https_proxy=${HTTPS_PROXY} --build-arg no_proxy=${no_proxy}

ifndef VERSION
	VERSION = DEV
endif

GOFLAGS := -ldflags "$(LDFLAGS)"

## Download dependencies and the run unit test and build the binary
all: install test clean build dockerize package

## Clean the dist directory
clean:
	@rm -rf $(BUILD_DEST)

ginkgo: generate
	@./scripts/ginkgo.coverage.sh

## Run unit test
test: ginkgo

## download dependencies to run this project
install:
	@which ginkgo > /dev/null || go get github.com/onsi/ginkgo/ginkgo
	@which gox > /dev/null || go get github.com/mitchellh/gox
	@which counterfeiter > /dev/null || go get github.com/maxbrunsfeld/counterfeiter
	@which dep > /dev/null || go get github.com/golang/dep/cmd/dep
	@which gocover-cobertura > /dev/null || go get github.com/t-yuki/gocover-cobertura
	dep ensure -vendor-only

## Run for local development
start:
	DATA_DIRECTORY="$$PWD/data" \
	go run cmd/mqttfaas/main.go

## Build the linux binary
build:
	@rm -rf $(BUILD_DEST)
	@mkdir -p $(BUILD_DEST) > /dev/null
	@CGO_ENABLED=false \
	gox \
	-arch="386" -arch="amd64" \
	-output "$(BUILD_DEST)/{{.Dir}}_{{.OS}}_{{.Arch}}" \
	-os="darwin linux windows" \
	$(GOFLAGS) \
	./cmd/mqttfaas/

## Build the docker image
dockerize: dockerize/lite

## Create DIND based image
dockerize/dind: #dockerize/lite
	docker build $(DOCKER_BUILD_ARGS) \
		--build-arg VERSION=${VERSION}-${GIT_COMMIT}-${GIT_DIRTY} \
		-t mqttfaas:dind -f Dockerfile.dind .

## Create lite image (Bring your docker runtime)
dockerize/lite:
	docker build $(DOCKER_BUILD_ARGS) \
		--build-arg VERSION=${VERSION}-${GIT_COMMIT}-${GIT_DIRTY} \
		-t mqttfaas:lite .

## Generate runs go:generate
generate:
	@git clean -xffd **/*fakes*
	@find . -iname interfaces.go | xargs -n 1 go generate

## Package Sample functions
package:
	./scripts/packagesamples.sh

## Prints the version info about the project
info:
	 @echo "Version:           ${VERSION}"
	 @echo "Git Commit:        ${GIT_COMMIT}"
	 @echo "Git Tree State:    ${GIT_DIRTY}"

## Print the dependency graph and open in MAC
dependencygraph:
	dep status -dot | dot -T png | open -f -a /Applications/Preview.app

## Prints this help command
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET}: ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
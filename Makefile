.DEFAULT_GOAL := dockerize
DOCKER_USERNAME := mqttfaas

ifdef TRAVIS_TAG
TAG = $(TRAVIS_TAG)
else
TAG := latest
endif

GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_DIRTY  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

CC ?=
BASE_IMAGE ?= 
ifeq (${ARCH},arm64)
CC := aarch64-linux-gnu-gcc
BASE_IMAGE := arm32v7/
endif

DOCKER_BUILD_ARGS=--build-arg http_proxy=${HTTP_PROXY} \
	--build-arg https_proxy=${HTTPS_PROXY} \
	--build-arg no_proxy=${no_proxy} \
	--build-arg VERSION=${GIT_COMMIT}-${GIT_DIRTY} \
	--build-arg GOARCH=${ARCH} \
	--build-arg CC=${CC} \
	--build-arg BASE_IMAGE=${BASE_IMAGE}

_dockerize:
	docker build $(DOCKER_BUILD_ARGS) \
		-t ${DOCKER_USERNAME}/${BINARY}:${ARCH}-${TAG} \
		-f cmd/${BINARY}/Dockerfile .

dockerize:
	docker pull openfaas/faas-swarm:0.5.0
	docker save -o faasswarm.tar.gz openfaas/faas-swarm:0.5.0
	@$(MAKE) dockerize-arch-amd64
	@## Commented out till I figure out a proper way to substitute the FROM Docker
	@# @$(MAKE) dockerize-arch-arm64

dockerize-arch-%:
	@$(MAKE) _dockerize ARCH=$* BINARY=imageloader
	@$(MAKE) _dockerize ARCH=$* BINARY=openfaas-runtime
	@$(MAKE) _dockerize ARCH=$* BINARY=httpinvoker

up: build
	docker-compose up --force-recreate --remove-orphans

cleanup:
	docker-compose stop
	docker-compose rm -f

docker/publish:
	# @$(MAKE) docker/publish-arm64
	@$(MAKE) docker/publish-amd64

docker/publish-%:
	@docker push ${DOCKER_USERNAME}/httpinvoker:$*-${TAG}
	@docker push ${DOCKER_USERNAME}/imageloader:$*-${TAG}
	@docker push ${DOCKER_USERNAME}/openfaas-runtime:$*-${TAG}
  
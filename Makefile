PROJECT_PREFIX = pipeline
GOLANG_VERSION = 1.15.3
ALPHINE_VERSION = 3.12
DOCKER_BUILD = docker build \
		-f docker/Dockerfile.${APP_LANG} \
		-t ${PROJECT_PREFIX}-${APP_NAME}:0.0.1 \
		--build-arg alphine_version=$(ALPHINE_VERSION) \
		--build-arg app_name=$(APP_NAME) \
		--build-arg golang_version=$(GOLANG_VERSION) \
		.

all: run build.all build.api_gateway
.PHONY: all

run:
	docker-compose up

build.all: build.api_gateway

build.api_gateway: APP_NAME := api_gateway
build.api_gateway: APP_LANG := golang
build.api_gateway:
	${DOCKER_BUILD}

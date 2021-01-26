PROJECT_PREFIX = pipeline
GOLANG_VERSION = 1.15.3
PYTHON_VERSION = 3.9.1
DOCKER_BUILD = docker build \
		-f docker/Dockerfile.${APP_LANG} \
		-t ${PROJECT_PREFIX}-${APP_NAME}:0.0.1 \
		--build-arg app_name=$(APP_NAME) \
		--build-arg golang_version=$(GOLANG_VERSION) \
		--build-arg python_version=$(PYTHON_VERSION) \
		.

all: run build.all build.apigateway build.foobar
.PHONY: all

run:
	docker-compose up

build.all: build.apigateway build.foobar

build.apigateway: APP_NAME := apigateway
build.apigateway: APP_LANG := golang
build.apigateway:
	${DOCKER_BUILD}

build.foobar: APP_NAME := foobar
build.foobar: APP_LANG := python
build.foobar:
	${DOCKER_BUILD}

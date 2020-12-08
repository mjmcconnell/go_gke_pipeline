PROJECT_PREFIX = pipeline
GOLANG_VERSION = 1.15.3
ALPHINE_VERSION = 3.12
DOCKER_BUILD = docker build \
		-f docker/Dockerfile.golang \
		-t ${PROJECT_PREFIX}-${APP_NAME}:0.0.1 \
		--build-arg alphine_version=$(ALPHINE_VERSION) \
		--build-arg app_name=$(APP_NAME) \
		--build-arg golang_version=$(GOLANG_VERSION) \
		.

all: build.all build.entrypoint build.enrichment build.response-builder
.PHONY: all

run:
	docker-compose up

build.all: build.entrypoint build.enrichment build.response-builder

build.entrypoint:
	$(eval APP_NAME := entrypoint)
	${DOCKER_BUILD}

build.enrichment:
	$(eval APP_NAME := enrichment)
	${DOCKER_BUILD}

build.response-builder:
	$(eval APP_NAME := response-builder)
	${DOCKER_BUILD}


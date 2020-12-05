DOCKER_RUN = docker run -it -v ${PWD}/${APP_DIR}:/app --workdir /app ${DOCKER_IMAGE}

all: run chef packer
.PHONY: all

run:
	docker-compose up

chef:
	$(eval APP_DIR := build/chef)
	$(eval DOCKER_IMAGE := chef/chefdk:4.9.17)
	$(DOCKER_RUN) berks vendor

packer:
	$(eval APP_DIR := build)
	$(eval DOCKER_IMAGE := hashicorp/packer:1.6.5)
	$(DOCKER_RUN) build /app/packer/docker-image.json

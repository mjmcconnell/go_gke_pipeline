run:
	docker-compose up

chef:
	docker run -it -v ${PWD}/build/chef:/app --workdir /app chef/chefdk:4.9.17 berks vendor

packer:
	docker run -it -v ${PWD}/build:/app --workdir /app hashicorp/packer:1.6.5 build /app/packer/docker-image.json

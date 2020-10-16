# Go GKE Pipeline

This repo is setup to play about with running Golang base services within a GKE environment.

## Setup

This project is setup with Docker (docker-compose), and Makefiles, for ease of use.

Simply make sure you have [docker-compose](https://docs.docker.com/compose/install/) installed locally.

### Minikube (Optional)

This repo not only contains the Golang source, but also comes with helm chart for installing the application on kubernetes.

To run application locally within a kubernetes environment, you can install [minikube](https://minikube.sigs.k8s.io/docs/start/).

## Local development

To get the application running locally, simply run:


    make docker.run

## Installing

### Minikube

    make mini.start
    make mini.install

### GKE

    make gke.install


## Useful links

* https://helm.sh/
* https://cloud.google.com/kubernetes-engine
* https://console.cloud.google.com/

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

## Kubernetes (GKE)

The services defined in this repo, are intended to be installed within a Kubernetes cluster (specifically GKE).

The proxy service is setup with session affinity, so that it is able to re-use connections that a given user has created.
To enable session affinity, the cluster needs to be setup with IP Alias (VPC-native), and the ingress should be configured to use NEG's.

See [here](https://cloud.google.com/kubernetes-engine/docs/how-to/ingress-features#session_affinity), for the required config, to enable session affinity in GKE.

### Installing

#### Minikube

    make mini.start
    make mini.install

#### GKE

    make gke.install


## Useful links

* https://helm.sh/
* https://cloud.google.com/kubernetes-engine
* https://console.cloud.google.com/

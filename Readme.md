# Go K8s Tracking Pipeline

This repo is setup to play about with running Golang base services within a kubernetes environment.

One of the main objects, is to be able to effectively track a request as it goes through different services within the k8s cluster.

Tracking is useful for keeping processing requests in a non blocking format, and allow to see the latency of the request at each step of its journey, and the overall request time.

## Infrastructure
* Chef -> Secure base image
* Packer -> Docker Image
* Terraform -> Kubernetes
* [Helm](https://helm.sh/) -> Kubernets resources
* Vault -> Secrets

## CI/CD
* Github CI

## Pipeline:
* Lambda style applications that enrich data
* Each app records heart beat of request as it passes through

## Monitoring:
* Prometheus
* Grafana
* [Jaeger](https://www.jaegertracing.io)

## Datastore:
* Redis (with keyspace event notifications)

## Request flow:
-> client request
-> entrypoint
    Adds meta data to the request
-> enrichment
    Adds enriches the user data
-> build response
    Builds the response object for the client
-> client response


## Services
### Entrypoint API
* Create tracking span, and add it to the client request
* Create datastore record
* Pass client request to the enrichment API
* Wait for record to be finalised
  * Redirect loopg
* Return response to client

### Enrichment API
* Pull data from the datastore, and enrich the user provided data

### Response Builder
* Build a response object based of the clients request

## Setup

This project is setup with Docker (docker-compose), and Makefiles, for ease of use.

Simply make sure you have [docker-compose](https://docs.docker.com/compose/install/) installed locally.

### Minikube (Optional)

This repo not only contains the Golang source, but also comes with helm chart for installing the application on kubernetes.

To run application locally within a kubernetes environment, you can install [minikube](https://minikube.sigs.k8s.io/docs/start/).

## Local development

To get the application running locally, simply run:

    make docker.run

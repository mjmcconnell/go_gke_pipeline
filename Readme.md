# Go K8s Tracking Pipeline

This repo is setup to play about with running Golang base services within a kubernetes environment.

One of the main objects, is to be able to effectively track a request as it goes through different services within the k8s cluster.

Tracking is useful for keeping processing requests in a non blocking format, and allow to see the latency of the request at each step of its journey, and the overall request time.

## Infrastructure
* Terraform -> Kubernetes
* [Helm](https://helm.sh/) -> Kubernets resources
* Vault -> Secrets

## CI/CD
* Github CI

## Pipeline:
* Lambda style applications that enrich data as it flows through the system
* Each app records heart beat of request as it passes through

## Monitoring:
* Prometheus
* Grafana
* [Jaeger](https://www.jaegertracing.io)
  * https://github.com/jaegertracing/jaeger-operator

## Request flow:
* Client request
* API gateway
  * Validates request
    * Failing requests return an error to the client
  * Determine which service should handle query
  * Message is placed sent to a given PubSub topic
* Secondary service
  * Picks up message from PubSub subscription
  * Processes request
  * Publishes response back onto the PubSub, for the API gateway to consume
* API gateway
  * Watches for responses
  * Gets response
    * Either directly from the PubSub message or Redis
  * Response is sent back to the client

## Services

### API Gateway
* Create tracking span, and add it to the client request
* Create datastore record
* Pass client request to the enrichment API
* Wait for record to be finalised
  * Redirect loopg
* Return response to client

## Setup

This project is setup with Docker (docker-compose), and Makefiles, for ease of use.

Simply make sure you have [docker-compose](https://docs.docker.com/compose/install/) installed locally.

### Minikube (Optional)

This repo not only contains the Golang source, but also comes with helm chart for installing the application on kubernetes.

To run application locally within a kubernetes environment, you can install [minikube](https://minikube.sigs.k8s.io/docs/start/).

## Local development

To get the application running locally, simply run:

    make docker.run

## Packages

[Logrus](golang logrus vs zerolog) - Logging
[Open Telemetry](https://github.com/open-telemetry/opentelemetry-go) - Metrics and tracing

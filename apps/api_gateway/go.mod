module github.com/mjmcconnell/go_gke_pipeline/apps/api_gateway

go 1.15

require (
	github.com/canthefason/go-watcher v0.2.4 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/sirupsen/logrus v1.7.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.16.0
	go.opentelemetry.io/otel v0.16.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.16.0
	go.opentelemetry.io/otel/sdk v0.16.0
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
)

package monitoring

import (
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// InitTracer creates a new trace provider instance and registers it as global trace provider.
func InitTracer() func() {
	// Create and install Jaeger export pipeline.
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint("http://jaeger:14268/api/traces"),
		jaeger.WithProcess(jaeger.Process{ServiceName: "api-gateway"}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)
	if err != nil {
		log := GetLogger()
		log.Fatal(err)
	}
	return flush
}

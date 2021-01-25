package monitoring

import (
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	"github.com/mjmcconnell/go_gke_pipeline/apps/apigateway/pkg/config"
)

// InitTracer creates a new trace provider instance and registers it as global trace provider.
func InitTracer(settings config.Settings) func() {
	// Create and install Jaeger export pipeline.
	flush, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(fmt.Sprintf("http://%s:%s/api/traces", settings.TracingHost, settings.TracingPort)),
		jaeger.WithProcess(jaeger.Process{ServiceName: settings.AppName}),
		jaeger.WithSDK(&sdktrace.Config{DefaultSampler: sdktrace.AlwaysSample()}),
	)
	if err != nil {
		log := GetLogger()
		log.Fatal(err)
	}
	return flush
}

func GetTracer(settings config.Settings) trace.Tracer {
	return otel.Tracer(settings.AppName)
}

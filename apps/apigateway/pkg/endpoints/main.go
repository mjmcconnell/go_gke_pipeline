// This file contains all meta endpoints, used to discribe the service
package endpoints

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/label"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/mjmcconnell/go_gke_pipeline/apps/apigateway/pkg/config"
	"github.com/mjmcconnell/go_gke_pipeline/apps/apigateway/pkg/monitoring"
)

type MainHandler struct {
	Settings config.Settings
}

func (h MainHandler) Register(r *mux.Router) {
	r.HandleFunc("/", h.Root)
}

func (h MainHandler) Root(w http.ResponseWriter, r *http.Request) {
	t := monitoring.GetTracer(h.Settings)
	_, span := t.Start(r.Context(), "FooBar", oteltrace.WithAttributes(label.String("id", "1234")))
	defer span.End()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(``))
}

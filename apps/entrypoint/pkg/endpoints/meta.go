// This file contains all meta endpoints, used to discribe the service
package endpoints

import (
	"net/http"

	"github.com/gorilla/mux"
)

type MetaHandler struct{}

func (h MetaHandler) Register(r *mux.Router) {
	r.HandleFunc("/-/readiness", h.Readiness)
	r.HandleFunc("/-/liveness", h.Liveness)
}

func (h MetaHandler) Readiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Ready`))
}

func (h MetaHandler) Liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Live`))
}

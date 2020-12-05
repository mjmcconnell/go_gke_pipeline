package app

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mjmcconnell/go_gke_pipeline/apps/entrypoint/pkg/endpoints"
)

func Run() error {
	router := mux.NewRouter()

	endpoints.MetaHandler{}.Register(router)

	http.ListenAndServe(":8080", router)

	return nil
}

package app

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mjmcconnell/go_gke_pipeline/apps/entrypoint/pkg/endpoints"
)

func Run() error {

	// Start the private server in the background
	go func() { startPrivateServer() }()
	// Start the public server
	startPublicServer()

	return nil
}

func startPublicServer() error {
	router := mux.NewRouter()
	endpoints.MainHandler{}.Register(router)
	err := http.ListenAndServe(":8080", router)
	return err
}

func startPrivateServer() error {
	router := mux.NewRouter()
	endpoints.MetaHandler{}.Register(router)
	err := http.ListenAndServe(":8888", router)
	return err
}

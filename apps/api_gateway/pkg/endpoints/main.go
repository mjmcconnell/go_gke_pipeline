// This file contains all meta endpoints, used to discribe the service
package endpoints

import (
	"net/http"

	"github.com/gorilla/mux"
)

type MainHandler struct{}

func (h MainHandler) Register(r *mux.Router) {
	r.HandleFunc("/", h.Root)
}

func (h MainHandler) Root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(``))
}

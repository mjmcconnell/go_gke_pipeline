package app

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gorilla/mux"
)

var RoutedServerHost, _ = url.Parse("http://webserver_1:8081/")
var DefaultServerHost, _ = url.Parse("http://webserver_2:8082/")


func Run() error {
	router := mux.NewRouter()

	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`Hello world`))
	})

	reverseProxy := httputil.NewSingleHostReverseProxy(DefaultServerHost)
	reverseProxy.Director = routeTraffic
	router.PathPrefix("/r").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reverseProxy.ServeHTTP(w, r)
	})

	http.ListenAndServe(":8080", router)

	return nil
}

func routeTraffic(req *http.Request) {
	req.Header.Add("X-Forwarded-Host", req.Host)
	req.URL.Scheme = "http"

	if isWhitelistedPath(req) {
		req.URL.Host = RoutedServerHost.Host
	} else {
		req.URL.Host = DefaultServerHost.Host
	}
}

func isWhitelistedPath(req *http.Request) bool {
	switch true {
	case req.URL.Path == "/create" && req.Method == "GET":
		return true
	default:
		return false
	}
}

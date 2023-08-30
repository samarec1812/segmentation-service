package http

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	readTimeoutServer  = 10
	writeTimeoutServer = 10
)

func NewHTTPServer(port string) *http.Server {
	handler := mux.NewRouter()

	s := &http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  readTimeoutServer * time.Second,
		WriteTimeout: writeTimeoutServer * time.Second,
	}

	AppRouter(handler)

	return s
}

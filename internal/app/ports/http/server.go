package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/samarec1812/segmentation-service/internal/app/service"
	"golang.org/x/exp/slog"
	"net/http"
	"time"
)

const (
	readTimeoutServer  = 10
	writeTimeoutServer = 10
)

func NewHTTPServer(port string, logger *slog.Logger, a service.App) *http.Server {
	handler := chi.NewRouter()

	s := &http.Server{
		Addr:         port,
		Handler:      handler,
		ReadTimeout:  readTimeoutServer * time.Second,
		WriteTimeout: writeTimeoutServer * time.Second,
	}

	AppRouter(handler, logger, a)

	return s
}

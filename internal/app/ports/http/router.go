package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"

	"github.com/samarec1812/segmentation-service/internal/app/service"
)

func AppRouter(r *chi.Mux, logger *slog.Logger, a service.App) {
	r.Use(middleware.RequestID)
	r.Use(LoggerMiddleware(logger))
	r.Use(middleware.Recoverer)

	r.Post("/segment", createSegment(logger, a))
	//r.HandleFunc("/segment").Methods("DELETE")
	//r.HandleFunc("/user-segment").Methods("POST")
	//r.HandleFunc("/user-segment").Methods("GET")

}

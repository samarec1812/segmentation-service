package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

func AppRouter(r *chi.Mux, logger *slog.Logger, a App) {
	r.Use(middleware.RequestID)
	r.Use(LoggerMiddleware(logger))
	r.Use(middleware.Recoverer)

	r.Post("/slug", createNote(a))
	//r.HandleFunc("/slug").Methods("DELETE")
	//r.HandleFunc("/user-slug").Methods("POST")
	//r.HandleFunc("/user-slug").Methods("GET")

}

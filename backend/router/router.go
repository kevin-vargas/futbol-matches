package router

import (
	"net/http"

	"backend/handler"

	"github.com/go-chi/chi/v5"
)

func SetupDefaultRoutes(r *chi.Mux) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func SetupRoutes(r *chi.Mux, h handler.Hello) {
	r.Get("/hello/{word}", h.Handle)
}

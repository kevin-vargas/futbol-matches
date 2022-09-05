package router

import (
	"net/http"

	"backend/handler"
	"backend/middleware"

	"github.com/go-chi/chi/v5"
)

func SetupDefaultRoutes(r *chi.Mux, metricsHandler http.Handler) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Handle("/metrics", metricsHandler)
}

func SetupAuthRoutes(r *chi.Mux, ha handler.Auth) {
	r.Post("/signup", ha.SingUp)
	r.Post("/login", ha.Login)
}

func SetupHelloRoutes(r *chi.Mux, h handler.Hello, a middleware.Middleware) {
	r.Group(func(r chi.Router) {
		r.Use(a)
		r.Get("/hello/{word}", h.Handle)
	})
}

func SetupMatchCrudRoutes(r *chi.Mux, h handler.MatchHandler) {
	r.Get("/match/{id}", h.Get)
	r.Post("/match", h.Create)
	r.Patch("/match/{id}", h.Update)
	r.Delete("/match/{id}", h.Delete)
}

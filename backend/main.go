package main

import (
	"net/http"

	"backend/config"
	"backend/handler"
	"backend/middleware"
	"backend/pkg/metrics"
	"backend/router"
	"backend/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	metrics.Register()
	cfg := config.New()
	s := service.New()
	h := handler.New(s)
	r := chi.NewRouter()

	// global middlewares
	r.Use(middleware.CountRequest)

	router.SetupDefaultRoutes(r, metrics.NewHandler())
	router.SetupRoutes(r, h)
	err := http.ListenAndServe(cfg.App.Port, r)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"net/http"

	"backend/config"
	"backend/handler"
	"backend/router"
	"backend/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.New()
	s := service.New()
	h := handler.New(s)
	r := chi.NewRouter()
	mh := handler.NewMatchHandler()
	router.SetupDefaultRoutes(r)
	router.SetupRoutes(r, h)
	router.SetupMatchCrudRoutes(r, mh)
	err := http.ListenAndServe(cfg.App.Port, r)
	if err != nil {
		panic(err)
	}
}

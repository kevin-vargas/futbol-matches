package main

import (
	"net/http"

	"backend/config"
	"backend/handler"
	"backend/middleware"
	"backend/pkg/jwt"
	"backend/repository"
	"backend/router"
	"backend/service"
	"backend/service/encrypt"

	"github.com/go-chi/chi/v5"
)

func main() {
	// config
	cfg := config.New()
	// repositories
	ur := repository.NewUser()

	// services
	s := service.New()
	e := encrypt.New()
	j := jwt.New(cfg.JWT.Token, cfg.JWT.Duration)
	sa := service.NewAuth(ur, e, j)

	// handlers
	h := handler.New(s)
	ha := handler.NewAuth(sa)

	r := chi.NewRouter()
	router.SetupDefaultRoutes(r)
	router.SetupAuthRoutes(r, ha)
	router.SetupHelloRoutes(r, h, middleware.Auth(j))

	err := http.ListenAndServe(cfg.App.Port, r)
	if err != nil {
		panic(err)
	}
}

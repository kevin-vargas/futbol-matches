package main

import (
	"net/http"

	"backend/config"
	"backend/handler"
	"backend/middleware"
	"backend/pkg/jwt"
	"backend/pkg/metrics"
	ur "backend/repository/user"
	"backend/router"
	"backend/service"
	"backend/service/encrypt"
	us "backend/service/user"

	"github.com/go-chi/chi/v5"
)

func main() {
	// config
	metrics.Register()
	cfg := config.New()

	// repositories
	userRepository := ur.NewUserRepository()

	// services
	encryptService := encrypt.New()
	jwtService := jwt.New(cfg.JWT.Token, cfg.JWT.Duration)
	userService := us.NewUserService(userRepository, encryptService, jwtService)

	// handlers
	ha := handler.NewAuth(userService)

	r := chi.NewRouter()
	ms := service.NewMatchService()
	mh := handler.NewMatchHandler(ms)

	uh := handler.NewUserHandler(userService)

	// global middlewares
	r.Use(middleware.CountRequest)

	// global middlewares
	r.Use(middleware.CountRequest)

	router.SetupDefaultRoutes(r, metrics.NewHandler())
	router.SetupAuthRoutes(r, ha)
	router.SetupMatchCrudRoutes(r, mh)
	router.SetupUserCrudRoutes(r, uh)
	err := http.ListenAndServe(cfg.App.Port, r)
	if err != nil {
		panic(err)
	}
}

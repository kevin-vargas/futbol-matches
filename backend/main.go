package main

import (
	mr "backend/repository/match"
	"net/http"

	"backend/config"
	"backend/handler"
	"backend/middleware"
	"backend/pkg/jwt"
	"backend/pkg/metrics"
	ur "backend/repository/user"
	"backend/router"
	"backend/service/encrypt"
	ms "backend/service/match"
	us "backend/service/user"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func main() {
	// config
	metrics.Register()
	cfg := config.New()
	// repositories
	userRepository := ur.NewUserRepository()
	matchRepository := mr.NewMatchRepository()

	// services
	encryptService := encrypt.New()
	jwtService := jwt.New(cfg.JWT.Token, cfg.JWT.Duration)

	userService := us.NewUserService(userRepository, encryptService, jwtService)

	matchService := ms.NewMatchService(matchRepository)

	// handlers
	ha := handler.NewAuth(userService)

	r := chi.NewRouter()

	mh := handler.NewMatchHandler(matchService)

	uh := handler.NewUserHandler(userService)

	// global middlewares
	r.Use(middleware.CountRequest)

	// global middlewares
	r.Use(middleware.CountRequest)

	router.SetupDefaultRoutes(r, metrics.NewHandler())
	router.SetupAuthRoutes(r, ha)
	router.SetupMatchCrudRoutes(r, mh)

	router.SetupUserCrudRoutes(r, uh)
	err := http.ListenAndServe(cfg.App.Port, cors.AllowAll().Handler(r))

	if err != nil {
		panic(err)
	}
}

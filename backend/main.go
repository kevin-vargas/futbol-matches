package main

import (
	mr "backend/repository/match"
	"backend/service"
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
	r "backend/service/redis"
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
	metricStore := r.New(cfg.Redis)

	// services
	encryptService := encrypt.New()
	jwtService := jwt.New(cfg.JWT.Token, cfg.JWT.Duration)

	userService := us.NewUserService(userRepository, encryptService, jwtService)

	matchService := ms.NewMatchService(matchRepository)

	metricService := service.NewMetric(metricStore)
	// handlers
	ha := handler.NewAuth(userService)

	r := chi.NewRouter()

	mh := handler.NewMatchHandler(matchService, metricStore, userService)

	uh := handler.NewUserHandler(userService)

	meh := handler.NewMetric(metricService)
	// global middlewares
	r.Use(middleware.CountRequest)

	// global middlewares
	r.Use(middleware.CountRequest)

	router.SetupDefaultRoutes(r, metrics.NewHandler())
	router.SetupAuthRoutes(r, ha)
	router.SetupMatchCrudRoutes(r, mh)
	router.SetupUserCrudRoutes(r, uh)
	router.SetupMetricRoutes(r, meh)
	err := http.ListenAndServe(cfg.App.Port, cors.AllowAll().Handler(r))
	if err != nil {
		panic(err)
	}
}

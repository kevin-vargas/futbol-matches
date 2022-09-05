package config

import "time"

type App struct {
	Port string
}

type JWT struct {
	Token    string
	Duration time.Duration
}

type Config struct {
	App App
	JWT JWT
}

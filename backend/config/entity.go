package config

import "time"

type App struct {
	Port string
}

type JWT struct {
	Token    string
	Duration time.Duration
}

type DB struct {
	Host     string
	Port     string
	Database string
}

type Config struct {
	App App
	JWT JWT
	DB  DB
}

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
	Database      string
	ConnectionURI string
}

type Redis struct {
	URI       string
	Name      string
	Pass      *string
	Retention int64
}

type Config struct {
	App   App
	JWT   JWT
	DB    DB
	Redis Redis
}

package config

import "time"

const (
	port          = ":8080"
	jwt_duration  = 1 * time.Hour
	env_jwt_token = "JWT_TOKEN"
	db_host       = "localhost"
	db_port       = "27017"
	db_database   = "futbol-matches"
)

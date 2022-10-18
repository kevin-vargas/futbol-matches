package config

import "time"

const (
	port          = ":8080"
	jwt_duration  = 1 * time.Hour
	env_jwt_token = "JWT_TOKEN"
	db_port       = "27017"
	db_database   = "futbol-matches"
)

const (
	env_db_host     = "DB_HOST"
	env_db_username = "DB_USERNAME"
	env_db_password = "DB_PASSWORD"
)

package config

import "time"

const (
	port          = ":8080"
	jwt_duration  = 1 * time.Hour
	env_jwt_token = "JWT_TOKEN"

	db_database = "futbol-matches"
)

const (
	env_connection_uri = "DB_CONNECTION_URI"
	env_redis_uri      = "REDIS_URI"
	env_redis_pass     = "REDIS_PASS"
	env_redis_name     = "REDIS_NAME"
)

package config

import "time"

const (
	port          = ":8080"
	jwt_duration  = 1 * time.Hour
	env_jwt_token = "JWT_TOKEN"
)

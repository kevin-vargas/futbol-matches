package config

import (
	"os"
)

// TODO: change to env variables

func New() Config {
	redisPassEnv := os.Getenv(env_redis_pass)
	var redisPass *string
	if redisPassEnv != "" {
		redisPass = &redisPassEnv
	}
	retention := int64(7 * 24 * 60 * 60 * 1000)
	return Config{
		App: App{
			Port: port,
		},
		JWT: JWT{
			Token:    os.Getenv(env_jwt_token),
			Duration: jwt_duration,
		},
		DB: DB{
			Database:      db_database,
			ConnectionURI: os.Getenv(env_connection_uri),
		},
		Redis: Redis{
			URI:       os.Getenv(env_redis_uri),
			Name:      os.Getenv(env_redis_name),
			Pass:      redisPass,
			Retention: retention,
		},
	}
}

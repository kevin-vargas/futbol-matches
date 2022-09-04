package config

import "os"

// TODO: change to env variables

func New() Config {
	return Config{
		App: App{
			Port: port,
		},
		JWT: JWT{
			Token:    os.Getenv(env_jwt_token),
			Duration: jwt_duration,
		},
	}
}

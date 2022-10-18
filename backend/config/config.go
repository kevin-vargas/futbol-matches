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
		DB: DB{
			Host:     db_host,
			Port:     db_port,
			Database: db_database,
			User:     db_username,
			Pass:     db_password,
		},
	}
}

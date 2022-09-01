package config

func New() Config {
	return Config{
		App: App{
			Port: port,
		},
	}
}

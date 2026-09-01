package config

type Config struct {
	Host    string
	Port    string
	GinMode string
}

func Load() *Config {
	return &Config{
		Host:    "0.0.0.0",
		Port:    "8080",
		GinMode: "debug",
	}
}

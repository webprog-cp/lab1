package config

type Config struct {
	Host       string
	Port       string
	GinMode    string
	APIPrefix  string
	APIVersion string
}

func (c *Config) BasePath() string {
	return c.APIPrefix + "/" + c.APIVersion
}

func Load() *Config {
	return &Config{
		Host:       "0.0.0.0",
		Port:       "8080",
		GinMode:    "debug",
		APIPrefix:  "/api",
		APIVersion: "v1",
	}
}

package config

type Config struct {
	ServiceAddress string
}

func NewConfig() *Config {
	return &Config{
		ServiceAddress: ":8081",
	}
}

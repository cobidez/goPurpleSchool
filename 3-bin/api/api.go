package config

import "os"

type Config struct {
	key string
}

func NewConfig() *Config {
	return &Config{
		key: os.Getenv("API_KEY"),
	}
}

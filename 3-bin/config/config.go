package config

import "os"

type Config struct {
	Key string
}

func NewConfig() *Config {
	return &Config{
		key: os.Getenv("API_KEY"),
	}
}

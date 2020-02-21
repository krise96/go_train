package apiserver

import "gostartup/src/repositories"

type APIConfig struct {
	Port     string `toml:"port"`
	LogLevel string `toml:"loglevel"`
	Store    *repositories.Config
}

func CreateConfig() *APIConfig {
	return &APIConfig{
		Port:     ":8080",
		LogLevel: "debug",
		Store:    repositories.CreateConfig(),
	}
}

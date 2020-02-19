package apiserver

type APIConfig struct {
	Port     string `toml:"port"`
	LogLevel string `toml:"loglevel"`
}

func CreateConfig() *APIConfig {
	return &APIConfig{
		Port:     ":8080",
		LogLevel: "debug",
	}
}

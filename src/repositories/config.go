package repositories

type Config struct {
	DBURL string `toml:"database_url"`
}

func CreateConfig() *Config {
	return &Config{}
}

package config

type Config struct {
	DefaultOffset int
	DefaultLimit  int

	Path         string
	UserFileName string
}

func Load() Config {

	cfg := Config{}

	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10

	cfg.Path = "./data"
	cfg.UserFileName = "/user.json"

	return cfg
}

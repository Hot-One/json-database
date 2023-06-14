package config

type Config struct {
	DefaultOffset int
	DefaultLimit  int

	UserPath         string
	UserFileName string

	ProductPath string
	ProductFileName string
}

func Load() Config {

	cfg := Config{}

	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10

	cfg.UserPath = "./data"
	cfg.UserFileName = "/user.json"

	cfg.ProductPath = "./data"
	cfg.ProductFileName = "/product.json"

	return cfg
}

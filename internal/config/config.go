package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddr string `env:"SERVER_ADDRESS"`
	BaseURL    string `env:"BASE_URL"`
	FilePath   string `env:"FILE_STORAGE_PATH"`
}

func InitConfig() (*Config, error) {
	cfg := new(Config)

	flag.StringVar(&cfg.ServerAddr, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&cfg.BaseURL, "b", "http://localhost:8080", "base address of the resulting shortened URL")
	flag.StringVar(&cfg.FilePath, "f", "localDB.json", "dir of the storage")

	flag.Parse()

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

package config

import (
	"flag"
	"github.com/akmyrzza/go-musthave-shortener/internal/cerror"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddr   string `env:"SERVER_ADDRESS"`
	BaseURL      string `env:"BASE_URL"`
	FilePath     string `env:"FILE_STORAGE_PATH"`
	DatabasePath string `env:"DATABASE_DSN"`
}

func InitConfig() (*Config, error) {
	cfg := new(Config)

	//pgxSource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "mysecret", "postgresdb")

	flag.StringVar(&cfg.ServerAddr, "a", "localhost:8080", "address and port to run server")
	flag.StringVar(&cfg.BaseURL, "b", "http://localhost:8080", "base address of the resulting shortened URL")
	flag.StringVar(&cfg.FilePath, "f", "localDB.json", "dir of the storage")
	flag.StringVar(&cfg.DatabasePath, "d", "", "path of database")

	flag.Parse()

	if err := env.Parse(cfg); err != nil {
		return nil, cerror.ErrEnvParseConfig
	}

	return cfg, nil
}

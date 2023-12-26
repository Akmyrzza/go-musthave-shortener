package main

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/app"
	"github.com/akmyrzza/go-musthave-shortener/internal/config"
	"log"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("error: initializing config: %d", err)
	}

	if err := app.Run(cfg); err != nil {
		log.Fatalf("err: %d", err)
	}
}

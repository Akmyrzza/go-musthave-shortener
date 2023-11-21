package app

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/config"
	"github.com/akmyrzza/go-musthave-shortener/internal/repository/local"
	"github.com/akmyrzza/go-musthave-shortener/internal/server/handler"
	"github.com/akmyrzza/go-musthave-shortener/internal/server/router"
	"github.com/akmyrzza/go-musthave-shortener/internal/server/util/logger"
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	"log"
	"net/http"
)

func Run(cfg *config.Config) {

	newLogger := logger.InitLogger()
	newRepository := local.NewLocalRepository()
	newService := service.NewServiceURL(newRepository)
	newHandler := handler.NewHandler(newService, cfg.BaseURL)

	newServer := &http.Server{
		Handler: router.InitRouter(newHandler, newLogger),
		Addr:    cfg.ServerAddr,
	}

	if err := newServer.ListenAndServe(); err != nil {
		log.Fatalf("error, runnnig http server: %d", err)
	}
}

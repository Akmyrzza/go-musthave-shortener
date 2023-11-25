package app

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/config"
	"github.com/akmyrzza/go-musthave-shortener/internal/repository/local"
	"github.com/akmyrzza/go-musthave-shortener/internal/server/handler"
	"github.com/akmyrzza/go-musthave-shortener/internal/server/router"
	"github.com/akmyrzza/go-musthave-shortener/internal/server/util/logger"
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	"net/http"
)

func Run(cfg *config.Config) error {

	newLogger := logger.InitLogger()
	newRepository := local.NewLocalRepository(cfg.FilePath)
	newService := service.NewServiceURL(newRepository)
	newHandler := handler.NewHandler(newService, cfg.BaseURL)

	newServer := &http.Server{
		Handler: router.InitRouter(newHandler, newLogger),
		Addr:    cfg.ServerAddr,
	}

	if err := newServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

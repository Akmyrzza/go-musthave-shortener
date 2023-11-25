package app

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/config"
	"github.com/akmyrzza/go-musthave-shortener/internal/repository"
	"github.com/akmyrzza/go-musthave-shortener/internal/server/handler"
	"github.com/akmyrzza/go-musthave-shortener/internal/server/router"
	"github.com/akmyrzza/go-musthave-shortener/internal/server/util/logger"
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	"net/http"
)

func Run(cfg *config.Config) error {

	lg := logger.InitLogger()
	repo, err := repository.NewLocalRepository(cfg.FilePath)
	if err != nil {
		return err
	}

	srv := service.NewServiceURL(repo)
	hndlr := handler.NewHandler(srv, cfg.BaseURL)

	newServer := &http.Server{
		Handler: router.InitRouter(hndlr, lg),
		Addr:    cfg.ServerAddr,
	}

	if err := newServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

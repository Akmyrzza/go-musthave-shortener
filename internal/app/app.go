package app

import (
	"log"
	"net/http"

	"github.com/akmyrzza/go-musthave-shortener/internal/cerror"
	"github.com/akmyrzza/go-musthave-shortener/internal/config"
	"github.com/akmyrzza/go-musthave-shortener/internal/repository"
	"github.com/akmyrzza/go-musthave-shortener/internal/server/handler"
	"github.com/akmyrzza/go-musthave-shortener/internal/server/router"
	"github.com/akmyrzza/go-musthave-shortener/internal/server/util/logger"
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
)

func Run(cfg *config.Config) error {
	lg := logger.InitLogger()
	defer func() {
		if err := lg.Sync(); err != nil {
			log.Fatalf("error: syncing file: %d", err)
		}
	}()

	repo, err := repository.NewInMemory(cfg.FilePath)
	if err != nil {
		return cerror.ErrInMemoryRepo
	}

	srv := service.NewServiceURL(repo)
	hndlr := handler.NewHandler(srv, cfg.BaseURL)

	newServer := &http.Server{
		Handler: router.InitRouter(hndlr, lg),
		Addr:    cfg.ServerAddr,
	}

	if err := newServer.ListenAndServe(); err != nil {
		return cerror.ErrRunningServer
	}

	return nil
}

package app

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/repository/store"
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

	db, err := store.InitDatabase(cfg.DatabasePath)
	if err != nil {
		return err
	}

	repo, err := repository.NewRepo(cfg.FilePath)
	if err != nil {
		return cerror.ErrInMemoryRepo
	}

	srv := service.NewServiceURL(repo, db)
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

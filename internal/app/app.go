package app

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/handler"
	"github.com/akmyrzza/go-musthave-shortener/internal/repository/local"
	"github.com/akmyrzza/go-musthave-shortener/internal/router"
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	"log"
	"net/http"
)

func Run() {

	newRepository := local.NewLocalRepository()
	newService := service.NewServiceURL(newRepository)
	newHandler := handler.NewHandler(newService)
	newRouter := router.InitRouter(newHandler)

	if err := http.ListenAndServe("localhost:8080", newRouter); err != nil {
		log.Fatalf("error, runnnig http server: %d", err)
	}
}

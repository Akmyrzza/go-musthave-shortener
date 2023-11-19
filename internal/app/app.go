package app

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/handler"
	"github.com/akmyrzza/go-musthave-shortener/internal/repository/local"
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	"log"
	"net/http"
)

func Run() {

	newRepository := local.NewLocalRepository()
	newService := service.NewServiceURL(newRepository)
	newHandler := handler.NewHandler(newService)

	mux := http.NewServeMux()
	mux.HandleFunc("/", newHandler.HandleURL)

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatalf("error, runnnig http server: %d", err)
	}
}
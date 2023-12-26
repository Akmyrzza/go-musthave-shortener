package router

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/server/handler"
	"github.com/akmyrzza/go-musthave-shortener/internal/server/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter(h *handler.Handler, l *zap.Logger) *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery(), middleware.LoggingRequest(l), middleware.CompressRequest())

	router.POST("/", h.CreateShortURL)
	router.GET("/:id", h.GetOriginalURL)
	router.POST("/api/shorten", h.CreateIDJSON)
	router.GET("/ping", h.Ping)
	router.POST("/api/shorten/batch", h.CreateShortURLs)

	return router
}

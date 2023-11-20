package router

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/handler"
	"github.com/akmyrzza/go-musthave-shortener/internal/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter(h *handler.Handler, l *zap.Logger) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.LoggingRequest(l), gin.Recovery())

	router.POST("/", h.CreateID)
	router.GET("/:id", h.GetURL)

	return router
}

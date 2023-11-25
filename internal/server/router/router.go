package router

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/server/handler"
	"github.com/akmyrzza/go-musthave-shortener/internal/server/middleware"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitRouter(h *handler.Handler, l *zap.Logger) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.LoggingRequest(l), gin.Recovery())
	router.Use(middleware.CompressRequest(), gin.Recovery())

	router.POST("/", h.CreateID)
	router.GET("/:id", h.GetURL)
	router.POST("/api/shorten", h.CreateIDJSON)

	return router
}

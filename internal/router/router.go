package router

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.POST("/", h.CreateID)
	router.GET("/:id", h.GetURL)

	return router
}

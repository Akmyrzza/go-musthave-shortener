package handler

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Handler struct {
	Service service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateID(ctx *gin.Context) {
	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
		return
	}

	id := h.Service.CreateID(string(reqBody))
	resId := "http://" + ctx.Request.Host + ctx.Request.RequestURI + id

	ctx.Header("Content-Type", "text/plain")
	ctx.String(http.StatusCreated, resId)
}

func (h *Handler) GetURL(ctx *gin.Context) {
	id, exist := ctx.Params.Get("id")
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no id in params"})
		return
	}

	originalURL, ok := h.Service.GetURL(id)
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}

	ctx.Header("Location", originalURL)
	ctx.Header("Content-Type", "text/plain")
	ctx.JSON(http.StatusTemporaryRedirect, nil)
}

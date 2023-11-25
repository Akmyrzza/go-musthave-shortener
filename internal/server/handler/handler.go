package handler

import (
	"encoding/json"
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
)

type Handler struct {
	Service service.Service
	BaseURL string
}

func NewHandler(s service.Service, BaseURL string) *Handler {
	return &Handler{
		Service: s,
		BaseURL: BaseURL,
	}
}

func (h *Handler) CreateID(ctx *gin.Context) {
	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
		return
	}

	id := h.Service.CreateID(string(reqBody))
	resultString := h.BaseURL + "/" + id

	ctx.Header("Content-Type", "text/plain")
	ctx.String(http.StatusCreated, resultString)
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

func (h *Handler) CreateIDJSON(ctx *gin.Context) {
	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
		return
	}

	stURL := struct {
		URL string `json:"url"`
	}{}

	if err = json.Unmarshal(reqBody, &stURL); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
		return
	}

	id := h.Service.CreateID(stURL.URL)
	resultString, err := url.JoinPath(h.BaseURL, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "url joining"})
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{"result": resultString})
}

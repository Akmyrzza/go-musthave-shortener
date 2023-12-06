package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateID(originalURL string) (string, error)
	GetURL(id string) (string, bool)
	Ping() error
}

type Handler struct {
	Service Service
	BaseURL string
}

func NewHandler(s Service, baseURL string) *Handler {
	return &Handler{
		Service: s,
		BaseURL: baseURL,
	}
}

func (h *Handler) CreateID(ctx *gin.Context) {
	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
		return
	}

	id, err := h.Service.CreateID(string(reqBody))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "creating id error"})
	}

	resultString, err := url.JoinPath(h.BaseURL, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "result path error"})
	}

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

	id, err := h.Service.CreateID(stURL.URL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "creating id error"})
	}

	resultString, err := url.JoinPath(h.BaseURL, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "url joining"})
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{"result": resultString})
}

func (h *Handler) Ping(ctx *gin.Context) {
	err := h.Service.Ping()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	ctx.JSON(http.StatusOK, "")
}

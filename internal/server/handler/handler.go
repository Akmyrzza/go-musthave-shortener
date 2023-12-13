package handler

import (
	"encoding/json"
	"github.com/akmyrzza/go-musthave-shortener/internal/model"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type ServiceURL interface {
	CreateShortURL(userID, originalURL string) (string, bool, error)
	GetOriginalURL(shortURL string) (string, error)
	Ping() error
	CreateShortURLs(userID string, urls []model.ReqURL) ([]model.ReqURL, error)
	GetAllURLs(userID string) ([]model.ResURL, error)
}

type Handler struct {
	Service ServiceURL
	BaseURL string
}

func NewHandler(s ServiceURL, b string) *Handler {
	return &Handler{
		Service: s,
		BaseURL: b,
	}
}

func (h *Handler) CreateShortURL(ctx *gin.Context) {
	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
		return
	}

	userID, err := ctx.Cookie("user_id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	id, exist, err := h.Service.CreateShortURL(userID, string(reqBody))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "creating id error"})
		return
	}

	resultString, err := url.JoinPath(h.BaseURL, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "result path error"})
		return
	}

	ctx.Header("Content-Type", "text/plain")
	if exist {
		ctx.String(http.StatusConflict, resultString)
		return
	}
	ctx.String(http.StatusCreated, resultString)
}

func (h *Handler) GetOriginalURL(ctx *gin.Context) {
	id, exist := ctx.Params.Get("id")
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no id in params"})
		return
	}

	originalURL, err := h.Service.GetOriginalURL(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
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

	userID, err := ctx.Cookie("user_id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	id, exist, err := h.Service.CreateShortURL(userID, stURL.URL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "creating id error"})
		return
	}

	resultString, err := url.JoinPath(h.BaseURL, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "url joining"})
		return
	}

	ctx.Header("Content-Type", "application/json")
	if exist {
		ctx.JSON(http.StatusConflict, gin.H{"result": resultString})
		return
	}
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

func (h *Handler) CreateShortURLs(ctx *gin.Context) {
	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
		return
	}

	var tmpURLs []model.ReqURL

	if err = json.Unmarshal(reqBody, &tmpURLs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
		return
	}

	userID, err := ctx.Cookie("user_id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	tmpURLs, err = h.Service.CreateShortURLs(userID, tmpURLs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "creating id error"})
		return
	}

	for i, v := range tmpURLs {
		resultString, err := url.JoinPath(h.BaseURL, v.ShortURL)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "creating id error"})
			return
		}
		tmpURLs[i].ShortURL = resultString
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, tmpURLs)
}

func (h *Handler) GetAllURLs(ctx *gin.Context) {
	userID, err := ctx.Cookie("user_id")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	data, err := h.Service.GetAllURLs(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	if len(data) == 0 {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	for i, v := range data {
		data[i].ShortURL, err = url.JoinPath(h.BaseURL, v.ShortURL)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, data)
}

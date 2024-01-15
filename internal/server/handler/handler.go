package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/akmyrzza/go-musthave-shortener/internal/cerror"
	"github.com/akmyrzza/go-musthave-shortener/internal/model"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

type ServiceURL interface {
	CreateShortURL(ctx context.Context, originalURL string) (string, error)
	GetOriginalURL(ctx context.Context, shortURL string) (string, error)
	Ping(ctx context.Context) error
	CreateShortURLs(ctx context.Context, urls []model.ReqURL) ([]model.ReqURL, error)
	GetAllURLs(ctx context.Context, userID string) ([]model.UserData, error)
	DeleteURLs(ctx context.Context, userID string, data []string)
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
	userID, exists := ctx.Get("userID")
	if !exists {
		userID = ""
	}

	user, ok := userID.(string)
	if !ok {
		log.Println("error create short url : type assertion error")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	newContext := context.WithValue(ctx.Request.Context(), model.KeyUserID("userID"), user)

	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
		return
	}

	id, cerr := h.Service.CreateShortURL(newContext, string(reqBody))
	if cerr != nil {
		if cerr != cerror.ErrAlreadyExist {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "creating id error"})
			return
		}
	}

	resultString, err := url.JoinPath(h.BaseURL, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "result path error"})
		return
	}

	ctx.Header("Content-Type", "text/plain")
	if cerr != nil {
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

	originalURL, err := h.Service.GetOriginalURL(ctx.Request.Context(), id)
	if err != nil {
		if errors.Is(err, cerror.ErrIsDeleted) {
			ctx.AbortWithStatus(http.StatusGone)
		}
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

	id, cerr := h.Service.CreateShortURL(ctx.Request.Context(), stURL.URL)
	if cerr != nil {
		if cerr != cerror.ErrAlreadyExist {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "creating id error"})
			return
		}
	}

	resultString, err := url.JoinPath(h.BaseURL, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "url joining"})
		return
	}

	ctx.Header("Content-Type", "application/json")
	if cerr != nil {
		ctx.JSON(http.StatusConflict, gin.H{"result": resultString})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"result": resultString})
}

func (h *Handler) Ping(ctx *gin.Context) {
	err := h.Service.Ping(ctx.Request.Context())
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, "")
		return
	}

	ctx.JSON(http.StatusOK, "")
}

func (h *Handler) CreateShortURLs(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		userID = ""
	}
	userIDString, ok := userID.(string)
	if !ok {
		log.Println("error create short urls : type assertion error")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	newContext := context.WithValue(ctx.Request.Context(), model.KeyUserID("userID"), userIDString)

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

	tmpURLs, err = h.Service.CreateShortURLs(newContext, tmpURLs)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "creating id error"})
		return
	}

	for i, v := range tmpURLs {
		resultString, err := url.JoinPath(h.BaseURL, v.ShortURL)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "creating id error"})
			return
		}
		tmpURLs[i].ShortURL = resultString
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, tmpURLs)
}

func (h *Handler) GetAllURLs(ctx *gin.Context) {
	newUser, exists := ctx.Get("newUser")

	_, ok := newUser.(bool)
	if ok && exists {
		log.Println("error get all urls: unauthorized")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userID, _ := ctx.Get("userID")
	userIDString, ok := userID.(string)
	if !ok {
		log.Println("error get all urls : type assertion error")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	data, err := h.Service.GetAllURLs(ctx.Request.Context(), userIDString)
	if err != nil {
		log.Printf("get all urls error: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	if len(data) == 0 {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	for i, v := range data {
		resultString, err := url.JoinPath(h.BaseURL, v.ShortURL)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, fmt.Errorf("joining path : %w", err))
			return
		}
		data[i].ShortURL = resultString
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, data)
}

func (h *Handler) DeleteURLs(ctx *gin.Context) {
	newUser, exists := ctx.Get("newUser")
	if exists && newUser.(bool) {
		log.Println("error delete urls: unauthorized")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userID, _ := ctx.Get("userID")
	userIDString, ok := userID.(string)
	if !ok {
		log.Println("error delete urls: type assertion error")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
		return
	}

	var data []string
	if err = json.Unmarshal(reqBody, &data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "bad request body"})
		return
	}

	h.Service.DeleteURLs(context.Background(), userIDString, data)

	ctx.AbortWithStatus(http.StatusAccepted)
}

package handler

import (
	"bytes"
	"encoding/json"
	"github.com/akmyrzza/go-musthave-shortener/internal/model"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/akmyrzza/go-musthave-shortener/internal/repository"
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var RandLength int = 16

func TestHandler_CreateShortURL(t *testing.T) {
	testRepository, err := repository.NewRepo("")
	if err != nil {
		log.Fatalf("error in repo: %d", err)
	}

	testService := service.NewServiceURL(testRepository)
	testHandler := NewHandler(testService, "http://localhost:8080")

	testRouter := gin.Default()
	testRouter.POST("/", testHandler.CreateShortURL)

	type want struct {
		code        int
		contentType string
	}

	tests := []struct {
		name string
		url  string
		want want
	}{
		{
			name: "test #1",
			url:  "www.google.com",
			want: want{
				code:        201,
				contentType: "text/plain",
			},
		},
		{
			name: "test #2",
			url:  "www.yandex.com",
			want: want{
				code:        201,
				contentType: "text/plain",
			},
		},
		{
			name: "test #3",
			url:  "www.netflix.com",
			want: want{
				code:        201,
				contentType: "text/plain",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.url))
			recorder := httptest.NewRecorder()

			testRouter.ServeHTTP(recorder, request)

			result := recorder.Result()
			require.NoError(t, result.Body.Close())

			assert.Equal(t, test.want.code, result.StatusCode)
			assert.Equal(t, test.want.contentType, result.Header.Get("Content-Type"))
		})
	}
}

func TestHandler_GetOriginalURL(t *testing.T) {
	testRepository, err := repository.NewRepo("")
	if err != nil {
		log.Fatalf("error in repo: %d", err)
	}
	testService := service.NewServiceURL(testRepository)
	testHandler := NewHandler(testService, "http://localhost:8080")
	testRouter := gin.Default()

	testRouter.POST("/", testHandler.CreateShortURL)
	testRouter.GET("/:id", testHandler.GetOriginalURL)

	type want struct {
		code     int
		location string
	}

	tests := []struct {
		name string
		url  string
		want want
	}{
		{
			name: "test #1",
			url:  "www.google.com",
			want: want{
				code:     307,
				location: "www.google.com",
			},
		},
		{
			name: "test #2",
			url:  "www.yandex.com",
			want: want{
				code:     307,
				location: "www.yandex.com",
			},
		},
		{
			name: "test #3",
			url:  "www.netflix.com",
			want: want{
				code:     307,
				location: "www.netflix.com",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(test.url))
			recorder := httptest.NewRecorder()

			testRouter.ServeHTTP(recorder, request)

			result := recorder.Result()
			resBody, err := io.ReadAll(result.Body)

			require.NoError(t, result.Body.Close())
			require.NoError(t, err)

			id := strings.TrimPrefix(string(resBody), "http://localhost:8080"+"/")

			requestGet := httptest.NewRequest(http.MethodGet, "/"+id, nil)
			recorderGet := httptest.NewRecorder()

			testRouter.ServeHTTP(recorderGet, requestGet)

			resultGet := recorderGet.Result()

			require.NoError(t, resultGet.Body.Close())
			assert.Equal(t, test.want.code, resultGet.StatusCode)
			assert.Equal(t, test.want.location, resultGet.Header.Get("Location"))
		})
	}
}

func TestHandler_CreateIDJSON(t *testing.T) {
	testRepository, err := repository.NewRepo("")
	if err != nil {
		log.Fatalf("error in repo: %d", err)
	}
	testService := service.NewServiceURL(testRepository)
	testHandler := NewHandler(testService, "http://localhost:8080")

	testRouter := gin.Default()
	testRouter.POST("/api/shorten", testHandler.CreateIDJSON)

	type want struct {
		code        int
		contentType string
	}

	type reqBody struct {
		URL string `json:"url"`
	}

	tests := []struct {
		name string
		url  reqBody
		want want
	}{
		{
			name: "test #1",
			url:  reqBody{URL: "www.google.com"},
			want: want{
				code:        201,
				contentType: "application/json",
			},
		},
		{
			name: "test #2",
			url:  reqBody{URL: "www.yandex.com"},
			want: want{
				code:        201,
				contentType: "application/json",
			},
		},
		{
			name: "test #3",
			url:  reqBody{URL: "www.netflix.com"},
			want: want{
				code:        201,
				contentType: "application/json",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			reqBody, err := json.Marshal(test.url)
			if err != nil {
				log.Fatalf("error, request test body: %d", err)
			}

			request := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewBuffer(reqBody))
			recorder := httptest.NewRecorder()

			testRouter.ServeHTTP(recorder, request)

			result := recorder.Result()

			require.NoError(t, result.Body.Close())
			assert.Equal(t, test.want.code, result.StatusCode)
			assert.Equal(t, test.want.contentType, result.Header.Get("Content-Type"))
		})
	}
}

func TestHandler_CreateShortURLs(t *testing.T) {
	type want struct {
		code int
	}

	tests := []struct {
		want want
	}{
		{
			want: want{
				code: 201,
			},
		},
	}

	type Sample struct {
		Correlation_id string `json:"correlation_id"`
		Original_url   string `json:"original_url"`
	}

	for _, test := range tests {
		t.Run("test mock", func(t *testing.T) {
			samples := []Sample{
				{
					Correlation_id: "1",
					Original_url:   "www.google.com",
				},
				{
					Correlation_id: "2",
					Original_url:   "www.netflix.com",
				},
				{
					Correlation_id: "3",
					Original_url:   "www.mail.ru",
				},
			}

			jsonData, err := json.Marshal(samples)
			if err != nil {
				log.Println(err)
			}

			request := httptest.NewRequest(http.MethodPost, "/api/shorten/batch", bytes.NewBuffer(jsonData))
			request.Header.Set("Content-Type", "application/json")
			recorder := httptest.NewRecorder()

			r := gin.Default()
			r.POST("/api/shorten/batch", newHandlerTest)
			r.GET("/:id", GetOriginalURL)

			//h := http.HandlerFunc(newHandlerTest)
			//h.ServeHTTP(recorder, request)
			r.ServeHTTP(recorder, request)
			resp := recorder.Result()
			require.NoError(t, resp.Body.Close())
			assert.Equal(t, test.want.code, resp.StatusCode)
		})
	}
}

var tmpArray []model.ReqURL

func newHandlerTest(ctx *gin.Context) {
	reqBody, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "result path error"})
		return
	}

	if err = json.Unmarshal(reqBody, &tmpArray); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "result path error"})
		return
	}

	var tmpArrayCopy []model.ReqURL
	copy(tmpArrayCopy, tmpArray)
	for i, _ := range tmpArrayCopy {
		randURL := randString()
		tmpArrayCopy[i].ShortURL = randURL
		tmpArrayCopy[i].OriginalURL = ""
		tmpArray[i].ShortURL = randURL
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "result path error"})
	}

	jsonRes, err := json.Marshal(tmpArrayCopy)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "result path error"})
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, jsonRes)
}

func GetOriginalURL(ctx *gin.Context) {
	id, exist := ctx.Params.Get("id")
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no id in params"})
		return
	}

	for _, v := range tmpArray {
		if v.ShortURL == id {
			ctx.Header("Location", v.OriginalURL)
			ctx.Header("Content-Type", "text/plain")
			ctx.JSON(http.StatusTemporaryRedirect, nil)
		}
	}

	ctx.JSON(http.StatusBadRequest, nil)
	return
}

func randString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, RandLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

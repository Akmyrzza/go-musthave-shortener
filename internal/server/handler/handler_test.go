package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
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

func TestHandler_CreateID(t *testing.T) {
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

func TestHandler_GetURL(t *testing.T) {
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

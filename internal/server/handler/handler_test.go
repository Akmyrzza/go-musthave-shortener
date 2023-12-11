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

	type Samples struct {
		Correlation_id string `json:"correlation_id"`
		Original_url   string `json:"original_url"`
	}

	for _, test := range tests {
		t.Run("test mock", func(t *testing.T) {
			sample := Samples{
				Correlation_id: "1",
				Original_url:   "www.google.com",
			}

			jsonData, err := json.Marshal(sample)
			if err != nil {
				log.Println(err)
			}

			request := httptest.NewRequest(http.MethodPost, "/api/shorten/batch", bytes.NewBuffer(jsonData))
			request.Header.Set("Content-Type", "application/json")
			recorder := httptest.NewRecorder()

			h := http.HandlerFunc(newHandlerTest)
			h.ServeHTTP(recorder, request)

			resp := recorder.Result()
			require.NoError(t, resp.Body.Close())
			assert.Equal(t, test.want.code, resp.StatusCode)
		})
	}
}

func newHandlerTest(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var tmpURLs []model.ReqURL

	if err = json.Unmarshal(reqBody, &tmpURLs); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, _ := range tmpURLs {
		randURL := randString()
		tmpURLs[i].ShortURL = randURL
		tmpURLs[i].OriginalURL = ""
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	jsonRes, err := json.Marshal(tmpURLs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonRes)
}

func randString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, RandLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

package handler

import (
	"github.com/akmyrzza/go-musthave-shortener/internal/repository/local"
	"github.com/akmyrzza/go-musthave-shortener/internal/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_createID(t *testing.T) {

	testRepository := local.NewLocalRepository()
	testService := service.NewServiceURL(testRepository)
	testHandler := NewHandler(testService)

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
			testHandler.createID(recorder, request)

			result := recorder.Result()
			defer result.Body.Close()

			assert.Equal(t, test.want.code, result.StatusCode)
			assert.Equal(t, test.want.contentType, result.Header.Get("Content-Type"))
		})
	}
}

func TestHandler_getURL(t *testing.T) {
	testRepository := local.NewLocalRepository()
	testService := service.NewServiceURL(testRepository)
	testHandler := NewHandler(testService)

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
			testHandler.createID(recorder, request)

			result := recorder.Result()

			defer result.Body.Close()
			resBody, err := io.ReadAll(result.Body)

			require.NoError(t, err)

			id := strings.TrimPrefix(string(resBody), "http://example.com/")

			requestGet := httptest.NewRequest(http.MethodGet, "/"+id, nil)
			recorderGet := httptest.NewRecorder()
			testHandler.getURL(recorderGet, requestGet)

			resultGet := recorderGet.Result()
			defer resultGet.Body.Close()

			assert.Equal(t, test.want.code, resultGet.StatusCode)
			assert.Equal(t, test.want.location, resultGet.Header.Get("Location"))
		})
	}

}

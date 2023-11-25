package middleware

import (
	"bytes"
	"compress/gzip"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

type gzipWriter struct {
	gin.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func CompressRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if ctx.Request.Header.Get("Content-Type") == "application/json" ||
			ctx.Request.Header.Get("Content-Type") == "text/html" {

			acceptEncodings := ctx.Request.Header.Values("Accept-Encoding")

			if foundHeader(acceptEncodings) {
				compressWriter := gzip.NewWriter(ctx.Writer)
				defer compressWriter.Close()

				ctx.Header("Content-Encoding", "gzip")
				ctx.Writer = &gzipWriter{ctx.Writer, compressWriter}
			}
		}

		contentEncodings := ctx.Request.Header.Values("Content-Encoding")

		if foundHeader(contentEncodings) {
			compressReader, err := gzip.NewReader(ctx.Request.Body)
			if err != nil {
				log.Fatalf("error: new reader: %d", err)
				return
			}
			defer compressReader.Close()

			body, err := io.ReadAll(compressReader)
			if err != nil {
				log.Fatalf("error: read body: %d", err)
				return
			}

			ctx.Request.Body = io.NopCloser(bytes.NewReader(body))
			ctx.Request.ContentLength = int64(len(body))
		}
		ctx.Next()
	}
}

func foundHeader(content []string) bool {
	for _, v := range content {
		if v == "gzip" {
			return true
		}
	}

	return false
}

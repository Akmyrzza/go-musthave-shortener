package middleware

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"

	"github.com/akmyrzza/go-musthave-shortener/internal/cerror"

	"github.com/gin-gonic/gin"
)

type gzipWriter struct {
	gin.ResponseWriter
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	n, err := w.Writer.Write(b)
	if err != nil {
		return 0, cerror.ErrWriteByte
	}
	return n, nil
}

func CompressRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get("Content-Type") == "application/json" ||
			ctx.Request.Header.Get("Content-Type") == "text/html" {
			acceptEncodings := ctx.Request.Header.Values("Accept-Encoding")

			if foundHeader(acceptEncodings) {
				compressWriter := gzip.NewWriter(ctx.Writer)
				defer func() {
					if err := compressWriter.Close(); err != nil {
						log.Fatalf("error: compress: %d", err)
					}
				}()

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
			defer func() {
				if err := compressReader.Close(); err != nil {
					log.Fatalf("error: syncing file: %d", err)
				}
			}()

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

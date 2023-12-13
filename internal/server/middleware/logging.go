package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingRequest(logger *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		logger.Info("Request",
			zap.String("path", ctx.Request.URL.Path),
			zap.String("method", ctx.Request.Method),
			zap.Duration("duration", time.Since(start)),
		)

		logger.Info("Responese",
			zap.Int("statusCode", ctx.Writer.Status()),
			zap.Int64("contentLength", int64(ctx.Writer.Size())),
		)
	}
}

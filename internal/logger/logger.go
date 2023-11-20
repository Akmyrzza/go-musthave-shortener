package logger

import (
	"go.uber.org/zap"
	"log"
)

func InitLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("error: initializing logger: %d", err)
	}

	defer logger.Sync()

	return logger
}

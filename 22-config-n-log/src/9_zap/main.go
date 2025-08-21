package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync()

	url := "test.test/best"
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.String("id", "3"),
		zap.Duration("backoff", time.Second),
	)
}

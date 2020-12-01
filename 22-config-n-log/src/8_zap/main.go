package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()
	defer logger.Sync() // flushes buffer, if any

	url := "test.test/best"
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

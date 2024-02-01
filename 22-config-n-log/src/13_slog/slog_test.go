package main

import (
	"os"
	"testing"

	"golang.org/x/exp/slog"
)

func TestSlog(t *testing.T) {
	slog.Info("Started", "count", 3)
	//logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	logger.Warn("aaaa!", "boom?", true)

	nestedLogger := logger.With("request_id", 42)

	processMessage(nestedLogger, "message to process")
}

func processMessage(nestedLogger *slog.Logger, s string) {
	nestedLogger.Info("start processing", "message", s)
}

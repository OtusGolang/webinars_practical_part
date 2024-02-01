package main

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"
)

func main() {
	//demoDefaultLogger()
	demoLoggers()
}

func demoDefaultLogger() {
	// default logger
	//ctx := context.Background()
	slog.Debug("debug 1", "count", 3)
	slog.Info("info 1", slog.Int("count", 3), "hi", "there")
	slog.Error("oh oh", "error", errors.New("something bad happen"))

	// can change default logger
	logConfig := &slog.HandlerOptions{
		AddSource:   false,
		Level:       slog.LevelDebug,
		ReplaceAttr: nil,
	}
	//logHandler := slog.NewJSONHandler(os.Stderr, logConfig)
	logHandler := slog.NewTextHandler(os.Stderr, logConfig)

	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	slog.Debug("debug 2", "count", 3)
	slog.Info("info 2", "count", 3)
}

func demoLoggers() {
	// setup:
	file, err := os.Create("main.log")
	if err != nil {
		panic(fmt.Errorf("create log file: %w", err))
	}
	defer file.Close()

	writer := io.MultiWriter(file, os.Stderr)
	logger := slog.New(slog.NewTextHandler(writer, nil))

	logger.Info("starting processing")

	// examples:
	requestLogger := logger.With("request_id", "12345")
	requestLogger.Info("going to call handler")
	err = handleRequest(requestLogger)
	//err = processRequest(requestLogger.WithGroup("handler")) // grouping is possible
	if err != nil {
		logger.Error("error while process request", "error", err)
	}
	logger.Info("processing finished")
}

func handleRequest(logger *slog.Logger) error {
	logger.Info("Processing started")
	logger.Warn("Processing took too long", slog.Duration("duration", 12*time.Second))
	return fmt.Errorf("some error")
}

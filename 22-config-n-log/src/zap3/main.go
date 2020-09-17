package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	url := "test.test/best"
	slogger := logger.Sugar()
	slogger.Infof("failed to fetch %s", url)

	plain := slogger.Desugar()
	plain.DPanic("ending message", zap.String("msg", "this is the end"))
}

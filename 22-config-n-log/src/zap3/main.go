package main

import (
	"go.uber.org/zap"
)

func main() {
	url := "test.test/best"
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	slogger := logger.Sugar()
	slogger.Infof("failed to fetch %s", url)
	plain := slogger.Desugar()
	plain.DPanic("ending message", zap.String("msg", "this is the end"))
}

package main

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = syslogTimeEncoder
	cfg.EncoderConfig.EncodeLevel = customLevelEncoder

	logger, _ := cfg.Build()

	logger.Info("This should have a syslog style timestamp")
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan 2 15:04:05"))
}

func customLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.String() + "]")
}

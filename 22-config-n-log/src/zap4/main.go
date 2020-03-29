package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func main() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = SyslogTimeEncoder
	cfg.EncoderConfig.EncodeLevel = CustomLevelEncoder
	logger, _ := cfg.Build()
	logger.Info("This should have a syslog style timestamp")
}
func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan 2 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.String() + "]")
}

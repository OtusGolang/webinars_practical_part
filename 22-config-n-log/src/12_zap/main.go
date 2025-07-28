package main

import (
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level       zapcore.Level `default:"debug"`
	LogEncoding string        `required:"true"`
}

func main() {
	appCfg := Config{}
	envconfig.MustProcess("MYAPP", &appCfg)

	logConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(appCfg.Level),
		DisableCaller:    true,
		Development:      true,
		Encoding:         appCfg.LogEncoding,
		OutputPaths:      []string{"stdout", "file_log.log"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
	}
	logConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // make output with colors!

	loggerRaw, err := logConfig.Build()
	if err != nil {
		panic(err)
	}
	logger := loggerRaw.Sugar()

	logger.Info("Started")
	logger.Debug("Debug mode enabled")

	// ---
	// logger.With("count", 1).Info("One banana")
	// logger.Desugar().With(zap.Int("count", 1)).Info("Two banana, no sugar")
	// logger.Infow("from infow", "count", 1, "name", "banana")

	// ---
	// parseResult, err := strconv.ParseBool("истина")
	// if err != nil {
	// 	logger.Errorw("error while parsing", "error", err)
	// }
	// logger.Info("parse result", "result", parseResult)

	// ---
	// requestLogger := logger.With("request_id", "12345")
	// processRequest(requestLogger)

	// ---
	// configuredLogger := logger.WithOptions(zap.AddStacktrace(zapcore.InfoLevel))
	// configuredLogger.Info("message example")

}

// func processRequest(logger *zap.SugaredLogger) {
// 	logger.Info("Processing started")
// 	logger.Error("Processing finished", "status", "error")
// }

// To show:
// 1. init, config
// 2. fields (sugar, desugar, infow)
// 3. logger inheritance

package main_test

import (
	"log"

	"go.uber.org/zap"
)

func ExampleLogger_example() {
	logger := zap.NewExample() // для тестов, пишет в stdout
	logger.With(zap.Int("field", 1)).Info("example message")

	// Output: {"level":"info","msg":"example message","field":1}
}

func ExampleLogger_development() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	logger.With(zap.Int("field", 1)).Info("example message")

	// Output: 2023-02-27T19:07:44.731-0500	INFO	log/example_test.go:23	example message	{"field": 1}
}

func ExampleLogger_production() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	logger.With(zap.Int("field", 1)).Info("example message")

	// Output: {"level":"info","ts":1677542766.6911411,"caller":"log/example_test.go:35","msg":"example message","field":1}
}

//Example and Production используют JSON Encoder, Development - Console
// logger.DPanic() паникует в Development, и не паникует - в Example / Production

// Development:
// 	Печатает stack trace для Warn и выше.
// 	Всегда печатает файл/номер строки
// 	Печатает уровень сообщения в верхнем регистре
// 	Печатает время в формате ISO8601 c миллисекундами
// Production:
// 	не логирует на уровне debug
// 	Добавляет stack trace в качестве поля на уровне Error/DPanic levels
// 	Добавляет caller в качестве поля
// 	Печатает timestamp в формате epoch
// 	Печатает уровень сообщения в нижнем регистре

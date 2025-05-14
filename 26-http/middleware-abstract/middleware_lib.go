package main

import (
	"log/slog"
	"strings"
	"time"
)

// ================= Примеры middleware =================

// Trimmer — изменяющий middleware: убирает пробелы по краям строки.
func Trimmer(next Processor) Processor {
	return func(input string) string {
		return next(strings.TrimSpace(input))
	}
}

// UpperCaser — ещё один изменяющий middleware: приводит строку к верхнему регистру.
func UpperCaser(next Processor) Processor {
	return func(input string) string {
		return next(strings.ToUpper(input))
	}
}

// Blocker — middleware, которое «коротко замыкает» цепочку: не вызывает next
// и сразу возвращает собственный результат.
func Blocker(next Processor) Processor {
	return func(input string) string {
		// Не передаём управление дальше
		return "[blocked] " + input
	}
}

// Logger — пример «сквозного» middleware: пишет лог перед и после обработки.
func Logger(tag string) Middleware {
	return func(next Processor) Processor {
		return func(input string) string {
			start := time.Now()
			slog.Info("start", "tag", tag, "input", input)
			out := next(input)
			slog.Info("done", "tag", tag, "output", out, "duration", time.Since(start))
			return out
		}
	}
}

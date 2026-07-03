package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/OtusGolang/webinars_practical_part/26-http/handler"
	"github.com/OtusGolang/webinars_practical_part/26-http/middleware"
	"github.com/lmittmann/tint"
)

// curl -d '{"candidate_id": 1, "passport": "test"}' -X POST 0.0.0.0:8080/vote
// curl 0.0.0.0:8080/stat
// curl 0.0.0.0:8080/stat/1

// powershell:
//  curl -uri http://localhost:8080/vote -method post -body '{"passport":"a", "candidate_id":123}'

func main() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stdout, nil)))

	h := handler.NewService()

	// Маршрутизация из Go 1.22+: префикс метода + точный путь vs поддерево + path-параметр.
	mux := http.NewServeMux()
	mux.HandleFunc("POST /vote", h.SubmitVote)             // только POST
	mux.HandleFunc("GET /stat", h.GetStats)                // точное совпадение → общая статистика
	mux.HandleFunc("GET /stat/{candidate_id}", h.GetStats) // path-параметр → статистика по кандидату
	mux.HandleFunc("GET /stat-stream", h.StatStream)       // websocket
	logger := middleware.NewLogger(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: logger,
	}

	// Контекст, отменяемый по SIGINT/SIGTERM.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		slog.Info("server start on", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("listen failed", "err", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	slog.Info("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("server shutdown failed", "err", err)
		return
	}
	slog.Info("server stopped gracefully")
}

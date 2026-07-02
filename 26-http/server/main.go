package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/OtusGolang/webinars_practical_part/26-http/handler"
	"github.com/lmittmann/tint"
)

type MyHandler struct {
	// some useful field
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serverHandler(w, r)
}

var vote = handler.NewService()

func serverHandler(w http.ResponseWriter, r *http.Request) {
	resp := &handler.Response{}
	switch r.URL.Path {
	case "/vote":
		vote.SubmitVote(w, r)
	case "/stat", "/stat/":
		vote.GetStats(w, r)
	default:
		resp.Error.Message = fmt.Sprintf("uri %s not found", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		handler.WriteResponse(w, resp)
	}
}

// curl -d '{"candidate_id": 1, "passport": "test"}' -X POST 0.0.0.0:8080/vote
// curl 0.0.0.0:8080/stat
// curl 0.0.0.0:8080/stat/?candidate_id=1
func main() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stdout, nil)))

	// var protocols http.Protocols        // http2 enable for non-https
	// protocols.SetUnencryptedHTTP2(true) // http2 enable for non-https
	handlerHttp := &MyHandler{}

	server := &http.Server{
		Addr: ":8080",
		//Handler:      http.HandlerFunc(serverHandler),
		Handler:      handlerHttp,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// Protocols:    &protocols, // http2 enable for non-https
	}

	// Контекст, который отменяется по SIGINT/SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Запускаем сервер в отдельной горутине, чтобы не блокировать main
	go func() {
		slog.Info("server start on", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("listen failed", "err", err)
			os.Exit(1)
		}
	}()

	// Ждём сигнал остановки
	<-ctx.Done()
	slog.Info("shutting down server...")

	// Даём активным запросам время завершиться
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		slog.Error("server shutdown failed", "err", err)
		return
	}
	slog.Info("server stopped gracefully")
}

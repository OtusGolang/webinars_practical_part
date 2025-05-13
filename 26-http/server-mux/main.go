package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/OtusGolang/webinars_practical_part/26-http/handler"
	"github.com/OtusGolang/webinars_practical_part/26-http/middleware"
	"github.com/lmittmann/tint"
)

// curl -d '{"candidate_id": 1, "passport": "test"}' -X POST 0.0.0.0:8080/vote
// curl 0.0.0.0:8080/stat
// curl 0.0.0.0:8080/stat/?candidate_id=1
func main() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stdout, nil)))

	h := handler.NewService()

	mux := http.NewServeMux()
	mux.HandleFunc("/vote", h.SubmitVote)
	mux.HandleFunc("/stat", h.GetStats)
	mux.HandleFunc("/stat/", middleware.IsArgExists(h.GetStats, "candidate_id"))
	// websocket handler
	mux.HandleFunc("/stat-stream", h.StatStream)

	logger := middleware.NewLogger(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: logger,
	}

	slog.Info("server start on", "addr", server.Addr)
	err := server.ListenAndServe()
	slog.Info("server stopped", "err", err)
}

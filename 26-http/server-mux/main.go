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

// powershell:
//  curl -uri http://localhost:8080/vote -method post -body '{"passport":"a", "candidate_id":123}'

func main() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stdout, nil)))

	h := handler.NewService()

	// ── Возможности маршрутизации ServeMux (Go 1.22+) ───────────────────────
	//
	// 1. Точное совпадение vs поддерево
	//    "/stat"   → только /stat
	//    "/stat/"  → совпадает с /stat/, /stat/foo, /stat/foo/bar (поддерево)
	//
	// 2. Префикс метода (Go 1.22+)
	//    "GET /vote"        → только GET-запросы
	//    "POST /vote"       → только POST-запросы
	//    "GET /stat/{id}"   → GET + параметр пути
	//
	// 3. Параметры пути (Go 1.22+)
	//    "GET /stat/{candidate_id}"         → r.PathValue("candidate_id")
	//    "GET /stat/{candidate_id}/details" → точный сегмент
	//    "GET /files/{path...}"             → wildcard: захватывает остаток пути
	//
	// 4. Маршрутизация по хосту
	//    mux.HandleFunc("admin.example.com/", adminHandler)
	//    mux.HandleFunc("example.com/", publicHandler)
	//
	// 5. Приоритет паттернов
	//    Более конкретный паттерн побеждает менее конкретный.
	//    Метод+путь важнее только пути: "GET /vote" > "/vote".
	//    Точное совпадение важнее поддерева: "/stat" > "/stat/".
	//    Конфликт (два одинаково конкретных паттерна) → panic при регистрации.
	//
	// 6. Редиректы
	//    Запрос "/stat" при наличии только "/stat/" → автоматический 301
	//    на "/stat/" (если "GET /stat" не зарегистрирован отдельно).
	//
	// Примеры:
	//    mux.HandleFunc("GET /vote", h.SubmitVote)
	//    mux.HandleFunc("POST /vote", h.SubmitVote)
	//    mux.HandleFunc("GET /stat/{candidate_id}", h.GetStats)  // r.PathValue("candidate_id")
	//    mux.HandleFunc("GET /files/{path...}", h.ServeFile)     // r.PathValue("path")
	// ────────────────────────────────────────────────────────────────────────

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

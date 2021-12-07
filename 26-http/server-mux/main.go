package main

import (
	"github.com/OtusGolang/webinars_practical_part/26-http/handler"
	"github.com/OtusGolang/webinars_practical_part/26-http/middleware"
	"log"
	"net/http"
)

// curl -d '{"candidate_id": 1, "passport": "test"}' -X POST 0.0.0.0:8080/vote
// curl 0.0.0.0:8080/stat
// curl 0.0.0.0:8080/stat/?candidate_id=1
func main() {
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

	log.Print("server start on port 8080")
	log.Fatal(server.ListenAndServe())
}

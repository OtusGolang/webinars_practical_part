package main

import (
	"io"
	"log/slog"
	"net/http"
)

func main() {

	resp, err := http.Get("http://127.0.0.1:8080/stat")
	if err != nil {
		slog.Error("error sending request", "err", err)
		return
	}
	defer resp.Body.Close() // <-- Зачем?

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("error reading response body", "err", err)
		return
	}

	slog.Info("responce body from stat", "body", body)
}

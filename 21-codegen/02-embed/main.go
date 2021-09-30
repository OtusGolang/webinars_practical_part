package main

import (
	_ "embed"
	"net/http"
)

type Handler struct {
}

//go:embed static/gopher.png
var gopherPngBytes []byte

func (*Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Write(gopherPngBytes)
}

func main() {
	if err := http.ListenAndServe(":8081", new(Handler)); err != nil {
		panic(err)
	}
}

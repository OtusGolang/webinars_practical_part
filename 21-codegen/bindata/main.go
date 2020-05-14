package main

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

func (*Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := Asset("static/gopher.png")
	if err != nil {
		fmt.Printf("%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(data)
	w.WriteHeader(http.StatusOK)
}

func main() {
	if err := http.ListenAndServe(":8081", new(Handler)); err != nil {
		panic(err)
	}
}

//go:generate go-bindata -o static_generated.go ./static

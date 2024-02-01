package main

import (
	"fmt"
	"net/http"
	//_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world!\n\n")
	})

	http.ListenAndServe(":8080", nil)
}

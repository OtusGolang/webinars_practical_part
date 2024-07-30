package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed static
var embedFiles embed.FS

func main() {
	// fsys, err := fs.Sub(embedFiles, "static")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	http.Handle("/", http.FileServer(http.FS(embedFiles)))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

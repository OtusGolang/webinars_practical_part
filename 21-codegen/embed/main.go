package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed static
var embedFiles embed.FS

func main() {
	fsys, err := fs.Sub(embedFiles, "static")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(fsys)))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal(err)
	}
}


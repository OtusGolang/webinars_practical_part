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

	// Для сравнения - можно раскомментировать и использовать реальную файловую систему вместо embed
	//real := os.DirFS(".")
	//http.Handle("/", http.FileServer(http.FS(real)))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}

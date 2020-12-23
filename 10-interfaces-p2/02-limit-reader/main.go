package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var r io.Reader

	r = strings.NewReader("hello")
	r = io.LimitReader(r, 4)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

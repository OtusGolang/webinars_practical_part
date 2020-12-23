package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	a := &App{
		Log: Logger{
			File: os.Stdout,
			Buf:  new(bytes.Buffer),
		},
	}
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}

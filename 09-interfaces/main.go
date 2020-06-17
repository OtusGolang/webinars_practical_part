package main

import (
	"log"
	"os"
)

func main() {
	a := &App{
		Log: Logger{
			File: os.Stdout,
		},
	}
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}

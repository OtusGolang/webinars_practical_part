package main

import (
	"flag"
	"fmt"
)

func main() {
	var msg string

	verbose := flag.Bool("verbose", false, "verbose output")
	flag.StringVar(&msg, "msg", "hello", "message to print")

	flag.Parse()

	if *verbose {
		fmt.Println("you say:", msg)
	} else {
		fmt.Println(msg)
	}
}

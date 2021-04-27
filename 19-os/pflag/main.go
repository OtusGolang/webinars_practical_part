package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

func main() {
	var msg string

	verbose := pflag.BoolP("verbose", "v", false, "verbose output")
	pflag.StringVarP(&msg, "msg", "m", "hello", "message to print")
	// pflag.Lookup("msg").NoOptDefVal = "bye"

	pflag.Parse()

	if *verbose {
		fmt.Println("you say:", msg)
	} else {
		fmt.Println(msg)
	}
}

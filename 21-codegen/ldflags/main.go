package main

import (
	"fmt"
)

var VersionString = "unset"

func main() {
	fmt.Println("Version:", VersionString)
}

// go run -ldflags '-X main.VersionString=1.0' .

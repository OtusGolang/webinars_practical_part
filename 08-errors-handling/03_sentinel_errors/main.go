package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello world!")

	buf := make([]byte, 5)
	if _, err := io.ReadFull(r, buf); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(buf))
	}

	longBuf := make([]byte, 64)
	if _, err := io.ReadFull(r, longBuf); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("ok")
	}
}

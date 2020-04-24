package main

import (
	"os"
)

func main() {
	file, err := os.Create("/tmp/fff")
	if err != nil {
		panic(err)
	}

	b := make([]byte, 1 << 20)
	_, err = file.Write(b)
	if err != nil {
		panic(err)
	}
	if err := file.Close(); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	tmpfile, err := ioutil.TempFile("/tmp", "example.")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tmpfile.Name())
	defer os.Remove(tmpfile.Name())

	content := []byte("temporary file's content\n")
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

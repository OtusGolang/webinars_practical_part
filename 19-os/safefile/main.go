package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dchest/safefile"
)

func main() {
	tmpfile, err := safefile.Create("/tmp/example.txt", 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tmpfile.Name())
	defer tmpfile.Close()

	content := []byte("temporary file's content\n")
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}

	time.Sleep(10 * time.Second)

	fmt.Println("commit")
	if err := tmpfile.Commit(); err != nil {
		log.Fatal(err)
	}
	time.Sleep(10 * time.Second)
}

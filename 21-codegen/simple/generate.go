// +build ignore

package main

import (
	"log"
	"os"
)

const helloFunc = `package main

import "fmt"

func Hello() {
	fmt.Println("Hello world!")
}`

func main() {
	f, err := os.OpenFile("code.go", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString(helloFunc)
	if err != nil {
		log.Fatal(err)
	}
}

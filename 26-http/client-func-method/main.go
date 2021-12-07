package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://127.0.0.1:8080/stat")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() // <-- Зачем?

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("responce body from stat: %s", body)
}

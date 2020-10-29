package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  int
	Meta json.RawMessage
}

func main() {
	data, err := json.Marshal(Person{
		Name: "Vasya",
		Age:  32,
		Meta: json.RawMessage(`{"settings": {}}`),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

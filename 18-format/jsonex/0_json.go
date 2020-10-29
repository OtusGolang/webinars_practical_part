package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string
	age  int
	Job  struct {
		Department string
		Title      string
	}
	Phones []string // Как сделать [] вместо null?
}

func main() {
	p1 := &Person{
		Name: "Vasya",
		age:  36,
		Job: struct {
			Department string
			Title      string
		}{Department: "Operations", Title: "Boss"},
	}

	j, err := json.MarshalIndent(p1, "--", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("p1 json %s\n", j)

	j, err = json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("p1 json %s\n", j)

	var p2 Person
	if err := json.Unmarshal(j, &p2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("p2: %v\n", p2)
}

// Как валидировать?

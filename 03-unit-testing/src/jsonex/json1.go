package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Job  struct {
		Department string
		Title      string
	}
}

func main() {
	p1 := &Person{
		Name: "Vasya",
		Age:  36,
		Job: struct {
			Department string
			Title      string
		}{Department: "Operations", Title: "Boss"},
	}

	j, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("p1 json %s\n", j)

	var p2 Person
	json.Unmarshal(j, &p2)
	fmt.Printf("p2: %+v\n", p2)
}

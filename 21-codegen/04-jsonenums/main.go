package main

import (
	"encoding/json"
	"fmt"
)

//go:generate jsonenums -type=Pill
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
)

func main() {
	p, err := json.Marshal(Aspirin)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(p))
}

package main

import (
	"fmt"
	"log"

	"github.com/vmihailenco/msgpack"
)

type Person struct {
	Name        string
	Surname     string
	Age         uint32
	ChildrenAge map[string]uint32
}

func main() {
	p := Person{
		Name:    "Ivan",
		Surname: "Remen",
		Age:     27,
	}
	p.ChildrenAge = make(map[string]uint32)
	p.ChildrenAge["Alex"] = 5
	p.ChildrenAge["Maria"] = 2

	marshaled, err := msgpack.Marshal(&p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nLength of marshaled: %v IMPL: %v\n", len(marshaled), string(marshaled))

	p2 := &Person{}
	if err := msgpack.Unmarshal(marshaled, p2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Unmarshled: %v\n", p2)
}

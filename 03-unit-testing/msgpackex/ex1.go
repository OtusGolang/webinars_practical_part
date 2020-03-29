package main

import (
	"fmt"

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
		Name:    "Rob",
		Surname: "Pike",
		Age:     27,
	}
	p.ChildrenAge = make(map[string]uint32)
	p.ChildrenAge["Alex"] = 5
	p.ChildrenAge["Maria"] = 2

	marshaled, _ := msgpack.Marshal(&p)

	fmt.Printf("Length of marshaled: %v\n", len(marshaled))
	fmt.Printf("Binary: %v\n", string(marshaled))

	p2 := &Person{}
	msgpack.Unmarshal(marshaled, p2)
	fmt.Printf("Unmarshled: %+v\n", p2)
}

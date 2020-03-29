package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func main() {
	p := &Person{
		Age:         27,
		Name:        "Rob",
		Surname:     "Pike",
		ChildrenAge: make(map[string]uint32),
	}
	p.ChildrenAge["Maria"] = 2
	p.ChildrenAge["Alex"] = 5

	marshaled, _ := proto.Marshal(p)

	fmt.Printf("Length of marshaled: %v\n", len(marshaled))
	fmt.Printf("Binary: %v\n", string(marshaled))

	p2 := &Person{}
	proto.Unmarshal(marshaled, p2)

	fmt.Printf("Unmarshaled %v\n", p2)
}

package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	p := &Person{
		Age:         27,
		Name:        "Ivan",
		Surname:     "Remen",
		ChildrenAge: make(map[string]uint32),
	}
	p.ChildrenAge["Maria"] = 2
	p.ChildrenAge["Alex"] = 5

	marshaled, _ := proto.Marshal(p)

	fmt.Printf("marshaled len %d message = %s\n", len(marshaled), string(marshaled))

	p2 := &Person{}
	proto.Unmarshal(marshaled, p2)

	fmt.Printf("Unmarshaled %v", p2)

}

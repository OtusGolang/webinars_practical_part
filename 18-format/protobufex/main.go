package main

import (
	"fmt"
	"log"

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

	marshaled, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("marshaled len %d message = %s\n", len(marshaled), string(marshaled))
	fmt.Println(marshaled)

	p2 := &Person{}
	if err := proto.Unmarshal(marshaled, p2); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nUnmarshaled %v", p2)
}

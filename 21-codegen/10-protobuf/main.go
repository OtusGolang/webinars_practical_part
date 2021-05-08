package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

//go:generate protoc --go_out=. Person.proto

func main() {
	p := new(Person)
	p.Name = "Anton"
	p.Mobile = append(p.Mobile, "8800553535")

	data, err := proto.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	p1 := Person{}
	err = proto.Unmarshal(data, &p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", p1)
}

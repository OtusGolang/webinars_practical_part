package intro

import (
	"fmt"
	"testing"
)

type Duck interface {
	Talk() string
	Walk()
	Swim()
}

type Dog struct {
	name string
}

func (d Dog) Talk() string {
	return "AGGGRRRR"
}

func (d Dog) Walk() {}

func (d Dog) Swim() {}

func quack(d Duck) {
	fmt.Println(d.Talk())
}

func TestInterfaces(t *testing.T) {
	quack(Dog{})
}

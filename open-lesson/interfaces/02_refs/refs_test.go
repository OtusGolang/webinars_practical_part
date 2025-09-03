package refs

import (
	"fmt"
	"testing"
)

type Speaker interface {
	SayHello()
}

type Human struct {
	Greeting string
}

func (h Human) SayHello() { // try reference receiver
	fmt.Println(h.Greeting)
}

func TestIfaceRefs(t *testing.T) {
	a := Human{Greeting: "Hello"}
	aref := &a

	var i1, i2 Speaker // interface

	i1 = a
	i2 = aref

	a.Greeting = "Updated text"

	a.SayHello()
	i1.SayHello()
	i2.SayHello()

	// {Updated text}
	// {Hello}
	// &{Updated text}

	after1 := i1.(Human)
	after2 := i2.(*Human)

	fmt.Println(after1, after2) // {Hello} &{Updated text}

	// i1: iface.data = copy
	// i2: iface.data = reference

}

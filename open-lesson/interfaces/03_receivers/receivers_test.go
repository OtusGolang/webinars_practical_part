package receivers

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

func (h *Human) SayHello() { // try reference receiver
	h.Greeting = h.Greeting + "!"
	fmt.Println(h.Greeting)
}

func TestReceivers(t *testing.T) {
	var s Speaker = &Human{Greeting: "Hello"}
	s.SayHello()
	s.SayHello()
	s = &Human{Greeting: "Hello2"}

	//s = &OtherHumanType{}
}

// TestUnaddressable shows another example of unaddressable values - map elements
func TestUnaddressable(t *testing.T) {
	exampleMap := map[string]Human{"a": {Greeting: "Hello"}}

	fmt.Println(exampleMap["a"])
	fmt.Println((exampleMap["a"])) // compiler: cannot take address of (exampleMap["a"])

	exampleVar := exampleMap["a"]
	fmt.Println(&exampleVar) // works
}

func TestReceivers_dark(t *testing.T) {
	// var s Speaker = Human{Greeting: "Hello"}
	// dark_magic(s)
	// s = Human{Greeting: "Hello2"}
	// dark_magic(s)
}

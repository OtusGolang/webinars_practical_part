package intro

import (
	"fmt"
)

type Speaker interface {
	SayHello()
}

type Human struct {
	Greeting string
}

func (h Human) SayHello() {
	fmt.Println(h.Greeting)
}

func main() {
	var s Speaker = Human{Greeting: "Hello"}
	s.SayHello()
}

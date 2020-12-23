package example

import "fmt"

type Greeter interface {
	hello() int
}

type Stranger interface {
	Bye() string
	Greeter
	fmt.Stringer
}

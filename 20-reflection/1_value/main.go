package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 42
	s := struct {
		string
		int
	}{"hello", 42}

	iv := reflect.ValueOf(i)
	sv := reflect.ValueOf(&s)

	fmt.Printf("%T: %v\n", iv, iv)
	fmt.Printf("%T: %v\n", sv, sv)
}

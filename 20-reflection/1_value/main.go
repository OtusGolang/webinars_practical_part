package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 42
	var s = struct {
		string
		int
	}{"hello", 42}

	iv := reflect.ValueOf(i)
	sv := reflect.ValueOf(&s)

	sv.Type()

	fmt.Printf("%T: %v\n", iv, iv)
	fmt.Printf("%T: %v\n\n", sv, sv)
}

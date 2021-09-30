package main

import (
	"fmt"
	"reflect"
)

func assertString(iv interface{}) (string, bool) {
	rv := reflect.ValueOf(iv)
	var s string
	var ok bool
	if rv.Kind() == reflect.String {
		s = rv.String()
		ok = true
	}
	return s, ok
}

func main() {
	var iv interface{} = "hello"
	s, ok := assertString(iv)
	fmt.Printf("%q %v\n", s, ok)

	s2, ok := assertString(42)
	fmt.Printf("%q %v\n", s2, ok)
}

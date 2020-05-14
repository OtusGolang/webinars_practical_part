package main

import (
	"fmt"
	"reflect"
)

func assertString(iv interface{}) (string, bool) {
	rv := reflect.ValueOf(iv)
	s := ""
	ok := false
	if rv.Kind() == reflect.String {
		s = rv.String()
		ok = true
	}
	return s, ok
}

func main() {
	var iv interface{} = "hello"
	s, ok := assertString(iv)
	fmt.Println(s, ok)
	s2, ok := assertString(42)
	fmt.Println(s2, ok)
}

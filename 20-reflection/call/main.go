package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := fmt.Printf
	v := reflect.ValueOf(f)
	args := []reflect.Value{
		reflect.ValueOf("test %d\n"),
		reflect.ValueOf(42),
	}
	ret := v.Call(args) // []reflect.Value
	fmt.Println(ret)
}

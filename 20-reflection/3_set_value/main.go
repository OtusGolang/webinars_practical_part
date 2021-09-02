package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 3.

	p := reflect.ValueOf(x)
	fmt.Println(p.Type(), ":", p.CanSet()) // ?

	p = reflect.ValueOf(&x)
	fmt.Println(p.Type(), ":", p.CanSet()) // ?

	v := p.Elem()
	fmt.Println(v.Type(), ":", v.CanSet()) // ?

	v.SetFloat(7.1)
	fmt.Println(x)
}

package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 3.4
	p := reflect.ValueOf(&x) // адрес переменной x
	v := p.Elem()            // переход по указателю
	fmt.Println(v.Type())    // float64
	fmt.Println(v.CanSet())  // true
	v.SetFloat(7.1)
	fmt.Println(x)
}

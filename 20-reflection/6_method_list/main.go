package main

import (
	"fmt"
	"reflect"
)

type Int int

func (i Int) Say() string {
	return "42"
}

func (i Int) Say50() string {
	return "50"
}

func main() {
	var obj Int
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i) // reflect.Value
		fmt.Println(v.Type().Method(i).Name, method)
	}

	sayMethod := v.MethodByName("Say") // reflect.Value
	fmt.Println(sayMethod.Call(nil))
}

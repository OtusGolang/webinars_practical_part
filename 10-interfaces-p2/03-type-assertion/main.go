package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string) // hello true
	fmt.Println(s, ok)

	r, ok := i.(fmt.Stringer) // <nil> false
	fmt.Println(r, ok)

	f, ok := i.(float64) // 0 false
	fmt.Println(f, ok)
}

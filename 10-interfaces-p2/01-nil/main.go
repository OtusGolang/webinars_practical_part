package main

import "fmt"

type I interface{ M() }

type T struct{}

func (T) M() {}

func main() {
	var t *T
	if t == nil {
		fmt.Println("t is nil")
	} else {
		fmt.Println("t is not nil")
	}
	var i I = t
	if i == nil {
		fmt.Println("i is nil")
	} else {
		fmt.Println("i is not nil")
	}
}

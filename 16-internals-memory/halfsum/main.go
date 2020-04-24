package main

import "fmt"

//go:noinline
func HalfSum(a, b int) int {
	c := a + b
	c /= 2
	return c
}

// go tool compile -S main.go
// go tool compile -m main.go
func main() {
	s := HalfSum(1, 2)
	fmt.Println(s)
}

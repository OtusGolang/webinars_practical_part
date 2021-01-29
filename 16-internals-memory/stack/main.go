package main

func main() {
	n := 4
	n2 := square(n)
	println(n2)
}

//go:noinline
func square(x int) int {
	return x * x
}

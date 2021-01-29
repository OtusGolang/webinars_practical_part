package main

func main() {
	n := answer()
	println(*n)
}

//go:noinline
func answer() *int {
	x := 42
	return &x
}

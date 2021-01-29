package main

func main() {
	n := 4
	square(&n)
	println(n)
}

//go:noinline
func square(x *int) {
	*x = (*x) * (*x)
}

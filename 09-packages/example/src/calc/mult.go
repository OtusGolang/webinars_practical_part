package calc

func Mult(args ...int) int {
	acc := 1
	for _, v := range args {
		acc *= v
	}

	return acc
}

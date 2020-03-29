package calc

func Sum(args ...int) int {
	acc := 0
	for _, v := range args {
		acc += v
	}

	return acc
}

package exbench

func Fast() int {
	acc := new(int)
	for i := 0; i < 10; i++ {
		acc2 := new(int)
		*acc2 = *acc + 1
		acc = acc2
	}

	return *acc
}

func Slow() int {
	acc := new(int)
	for i := 0; i < 1000; i++ {
		acc2 := new(int)
		*acc2 = *acc + 1
		acc = acc2
	}

	return *acc
}
package prealloc

func Fast() {
	const sz = 32000
	a := make([]int, 0)
	for i := 0; i < sz; i++ {
		a = append(a, i)
	}
}

func Slow() {
	const sz = 32000
	a := make([]int, 0, sz)
	for i := 0; i < sz; i++ {
		a = append(a, i)
	}
}


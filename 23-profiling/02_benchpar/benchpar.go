package benchpar

import (
	"sync"
)

var (
	mu = sync.Mutex{}
)

func Fast() int {
	acc := 0
	for i := 0; i < 1000; i++ {
		acc++
	}

	return acc
}

func Slow() int {
	mu.Lock()
	defer mu.Unlock()

	acc := 0
	for i := 0; i < 1000; i++ {
		acc++
	}

	return acc
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randFn := func() interface{} { return rand.Intn(100000000) }

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, randFn))
	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start)) // ~10s
}

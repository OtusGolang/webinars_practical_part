package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	randFn := func() interface{} { return rand.Intn(100000000) }

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	// rand -> repeatFn -> toInt 				fanIn -> take ->
	//						  -> primeFinder ->
	//						  -> primeFinder ->
	//						  -> primeFinder ->
	//						  -> primeFinder ->
	randIntStream := toInt(done, repeatFn(done, randFn))
	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)

	finders := make([]<-chan interface{}, numFinders)
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	fmt.Println("Primes:")
	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start)) // ~3s
}

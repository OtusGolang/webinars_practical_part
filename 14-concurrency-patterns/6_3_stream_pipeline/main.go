package main

import "fmt"

func main() {
	generator := func(integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, i := range integers {
				intStream <- i
			}
		}()
		return intStream
	}

	multiply := func(intStream <-chan int, multiplier int) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for i := range intStream {
				multipliedStream <- i * multiplier
			}
		}()
		return multipliedStream
	}

	add := func(intStream <-chan int, additive int) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for i := range intStream {
				addedStream <- i + additive
			}
		}()
		return addedStream
	}

	intStream := generator(1, 2, 3, 4)
	pipeline := multiply(add(multiply(intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}

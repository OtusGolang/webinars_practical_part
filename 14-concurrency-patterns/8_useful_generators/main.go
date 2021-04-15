package main

import (
	"fmt"
	"math/rand"
)

func main() {
	repeatFn := func(done <-chan struct{}, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	/*
		repeatBy := func(done <-chan struct{}, fn func() interface{}, interval time.Duration) <-chan interface{} {
			valueStream := make(chan interface{})
			go func() {
				t := time.NewTicker(interval)

				defer func() {
					t.Stop()
					close(valueStream)
				}()

				for {
					select {
					case <-done:
						return
					case <-t.C:
						select {
						case <-done:
						case valueStream <- fn():
						}
					}
				}
			}()
			return valueStream
		}
	*/

	take := func(done <-chan struct{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream: // Нет ли ничего подозрительного? :)
				}
			}
		}()
		return takeStream
	}

	done := make(chan struct{})
	defer close(done)

	randFn := func() interface{} { return rand.Int() }

	for num := range take(done, repeatFn(done, randFn), 10) {
		fmt.Println(num)
	}
}

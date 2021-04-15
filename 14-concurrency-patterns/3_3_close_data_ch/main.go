package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(strings <-chan int) <-chan struct{} {
		terminated := make(chan struct{})
		go func() {
			defer func() {
				fmt.Println("doWork exited.")
				close(terminated)
			}()

			for s := range strings {
				fmt.Println(s)
			}
		}()
		return terminated
	}

	data := make(chan int)
	terminated := doWork(data)

	go func() {
		for i := 0; i < 3; i++ {
			data <- i
			time.Sleep(1 * time.Second)
		}

		fmt.Println("Canceling doWork goroutine...")
		close(data)
	}()

	<-terminated
	fmt.Println("Done.")
}

package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer func() {
				fmt.Println("doWork exited.")
				close(completed)
			}()
			for s := range strings {
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)

	time.Sleep(time.Second * 5)
	fmt.Println("Done.")
}

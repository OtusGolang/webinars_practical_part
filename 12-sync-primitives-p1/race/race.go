package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	text := ""
	wg.Add(2)
	go func() {
		text = "hello world"
		wg.Done()
	}()
	go func() {
		fmt.Println(text)
		wg.Done()
	}()
	wg.Wait()
}

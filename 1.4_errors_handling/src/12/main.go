package main

import "fmt"

func server(workChan <-chan int) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work int) {
	defer fmt.Printf("work %d are done\n", work)
	do(work)
}

func do(work int) {
	fmt.Println("success")
	//panic("failed")
}

func main() {
	workChan := make(chan int)
	defer close(workChan)

	go server(workChan)

	for i := 0; i < 10; i++ {
		workChan <- i
	}
}

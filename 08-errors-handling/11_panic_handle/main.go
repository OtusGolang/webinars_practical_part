package main

import (
	"log/slog"
)

func server(workChan <-chan int) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work int) {
	defer func() {
		slog.Info("work is done", "work", work)

		// if err := recover(); err != nil {
		// 	slog.Info("work failed: ", "error", err)
		// }
	}()

	do(work)
}

func do(work int) {
	slog.Info("success", "work", work)
	// panic("failed")
}

func main() {
	workChan := make(chan int)
	defer close(workChan)

	go server(workChan)

	for i := 0; i < 10; i++ {
		workChan <- i
	}
}

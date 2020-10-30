package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("pid:", os.Getpid())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL)
	signal.Ignore(syscall.SIGTERM)

	for s := range c {
		fmt.Println("Got signal:", s)
	}
}

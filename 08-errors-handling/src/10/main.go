package main

import "fmt"

func process(r []int) {
	for _, rr := range r {
		defer fmt.Printf("finish processing for %d\n", rr)
		fmt.Printf("processing %d\n", rr)
	}
}

func main() {
	r := []int{1, 2}
	process(r)
}

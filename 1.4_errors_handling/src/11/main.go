package main

import (
	"fmt"
)

func calcSumStartingFrom5(r []int) int {
	defer func() {
		fmt.Printf("sum(%d elements)=%d\n", len(r), sum)
	}()

	sum := 0
	for _, rr := range r {
		sum += rr
	}

	return sum
}

func main() {
	r := []int{1, 2}
	calcSumStartingFrom5(r)

	r = []int{}
	calcSumStartingFrom5(r)
}

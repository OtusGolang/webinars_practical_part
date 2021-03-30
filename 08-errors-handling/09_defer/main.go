package main

import (
	"fmt"
)

func calcSum(r []int) int {
	count := 0
	defer func() {
		fmt.Printf("elements count: %d\n", count)
	}()

	sum := 0
	for _, rr := range r {
		sum += rr
		count++
	}

	return sum
}

func main() {
	r := []int{1, 2, 3, 4, 5}
	sum := calcSum(r)
	fmt.Println(sum)
}

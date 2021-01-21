package main

import "fmt"

type (
	TransformFn func(v, k int) int
	Stage       func(values []int) []int
)

func main() {
	multiplier := func(v, k int) int { return v * k }
	adder := func(v, k int) int { return v + k }

	newStage := func(fn TransformFn, n int) Stage {
		return func(values []int) []int {
			result := make([]int, len(values))
			for i, v := range values {
				result[i] = fn(v, n)
			}
			return result
		}
	}

	multiply := newStage(multiplier, 2)
	add := newStage(adder, 1)

	ints := []int{1, 2, 3, 4}
	for _, v := range add(multiply(ints)) {
		fmt.Println(v)
	}
}

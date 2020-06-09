package main

import (
	"fmt"
)

type IntStack struct {
	stack []int
}

func (s *IntStack) Push(v int) {
	s.stack = append(s.stack, v)
}

func (s *IntStack) Pop() int {
	if len(s.stack) == 0 {
		return 0
	}
	v := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return v
}

func main() {
	var s IntStack
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("expected 30, got %d\n", s.Pop())
	fmt.Printf("expected 20, got %d\n", s.Pop())
	fmt.Printf("expected 10, got %d\n", s.Pop())
	fmt.Printf("expected 0, got %d\n", s.Pop())
}

package variables

import (
	"fmt"
	"testing"
)

type S struct{}

func TestVariables(t *testing.T) {
	var s1 S
	fmt.Printf("1: %#v, %T\n", s1, s1)

	var s2 any

	s2 = 123
	fmt.Printf("2: %#v, %T\n", s2, s2)
	s2 = S{}
	fmt.Printf("3: %#v, %T\n", s2, s2)
}

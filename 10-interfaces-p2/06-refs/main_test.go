package refs

import (
	"fmt"
	"testing"
)

type A struct {
	F1 int
}

func TestIfaceRefs(t *testing.T) {
	a := A{F1: 42}
	aref := &a

	var i1, i2 any

	i1 = a
	i2 = aref

	a.F1 = 100500

	fmt.Println(a, i1, i2) // {100500} {42} &{100500}

	// i1: iface.data = copy
	// i2: iface.data = reference

}

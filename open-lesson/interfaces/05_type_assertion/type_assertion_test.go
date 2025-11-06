package typeassertion

import (
	"fmt"
	"testing"
)

func TestTypeAssertion(t *testing.T) {
	var i any = "hello" // any = interface{}

	r1 := i.(string)
	fmt.Println(r1)

	r2, ok := i.(string)
	fmt.Println(r2, ok)

	r3, ok := i.(fmt.Stringer)
	fmt.Println(r3, ok)

	r4, ok := i.(float64)
	fmt.Println(r4, ok)
}

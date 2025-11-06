package intro

import (
	"fmt"
	"testing"
)

func PrintAll(vals []any) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func TestAny(t *testing.T) {
	names := []any{"stanley", "david", "oscar"}

	PrintAll(names)
}

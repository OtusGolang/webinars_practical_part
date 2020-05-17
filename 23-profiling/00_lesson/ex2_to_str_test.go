package bench_test

import (
	"fmt"
	"testing"
)

func toString(v int) string {
	return fmt.Sprintf("%d", v)
	// return strconv.Itoa(v)
}

func BenchmarkToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		toString(42)
	}
}

package bench_test

import (
	"testing"
)

func ConcatString(m map[string]struct{}) string {
	a := ""
	for s := range m {
		a += s
	}
	return a
}

// func ConcatString(m map[string]struct{}) string {
// 	builder := strings.Builder{}
// 	for s := range m {
// 		builder.WriteString(s)
// 	}
// 	return builder.String()
// }

func BenchmarkConcatString(b *testing.B) {
	m := map[string]struct{}{
		"1": struct{}{},
		"2": struct{}{},
		"3": struct{}{},
		"4": struct{}{},
		"5": struct{}{},
	}

	for i := 0; i < b.N; i++ {
		ConcatString(m)
	}
}

package generics

import (
	"testing"
)

// Benchmark for the function with empty interfaces
func BenchmarkFunctionWithoutGenerics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertMapWithoutGenerics(map[string]int32{"1": 1, "2": 2})
	}
}

// Benchmark for the function with generics
func BenchmarkFunctionWithGenerics(b *testing.B) {
	convertMap := ConvertMapWithGenerics[string, int32, int64]
	for i := 0; i < b.N; i++ {
		convertMap(map[string]int32{"1": 1, "2": 2})
	}
}

func BenchmarkGmaxString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GMax("a", "b")
	}
}

func BenchmarkSmax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SMax("a", "b")
	}
}

func BenchmarkGmaxInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GMax(1, 2)
	}
}

func BenchmarkImax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IMax(1, 2)
	}
}

package reusehttp

import (
	"testing"
)

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fast()
	}
}

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slow()
	}
}

func BenchmarkParallelFast(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Fast()
		}
	})
}

func BenchmarkParallelSlow(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Slow()
		}
	})
}

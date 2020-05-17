package bench

import (
	"sync"
	"testing"
)

var mu sync.Mutex
var v int64

func Sum() {
	mu.Lock()
	v++
	mu.Unlock()
}

func BenchmarkConcurency(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Sum()
		}
	})
}

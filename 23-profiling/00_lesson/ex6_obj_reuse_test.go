package bench_test

import (
	"encoding/json"
	"testing"
)

type A struct {
	I int
}

func BenchmarkReuseObject(b *testing.B) {
	buf := []byte("{\"I\": 32}")
	for i := 0; i < b.N; i++ {
		a := &A{}
		json.Unmarshal(buf, a)
	}
}

// func BenchmarkReuseObject(b *testing.B) {
// 	buf := []byte("{\"I\": 32}")
// 	a := &A{}
// 	for i := 0; i < b.N; i++ {
// 		*a = A{}
// 		json.Unmarshal(buf, a)
// 	}
// }

// func BenchmarkReuseObject(b *testing.B) {
// 	p := sync.Pool{
// 		New: func() interface{} {
// 			return &A{}
// 		},
// 	}
// 	buf := []byte("{\"I\": 32}")
// 	for i := 0; i < b.N; i++ {
// 		a := p.Get().(*A)
// 		json.Unmarshal(buf, a)
// 		*a = A{}
// 		p.Put(a)
// 	}
// }

package bench_test

import (
	"encoding/json"
	"testing"
)

type A struct {
	I int
}

func BenchmarkReuseObject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := &A{}
		json.Unmarshal([]byte("{\"i\": 32}"), a)
	}
}

// func BenchmarkReuseObject(b *testing.B) {
// 	a := &A{}
// 	for i := 0; i < b.N; i++ {
// 		*a = A{}
// 		json.Unmarshal([]byte("{\"i\": 32}"), a)
// 	}
// }

// func BenchmarkReuseObject(b *testing.B) {
// 	p := sync.Pool{
// 		New: func() interface{} {
// 			return &A{}
// 		},
// 	}
// 	var a *A
// 	for i := 0; i < b.N; i++ {
// 		a = p.Get().(*A)
// 		json.Unmarshal([]byte("{\"i\": 32}"), a)
// 		p.Put(a)
// 	}
// }

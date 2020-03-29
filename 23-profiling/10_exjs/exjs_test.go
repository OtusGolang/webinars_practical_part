package exjs

import "testing"

func BenchmarkStandardMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardMarshal()
	}
}

func BenchmarkStandardUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardUnmarshal()
	}
}

func BenchmarkEasyMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EasyMarshal()
	}
}

func BenchmarkEasyUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EasyMarshal()
	}
}

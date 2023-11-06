package main

import (
	"testing"

	benchproto "benchs/proto"
	"github.com/golang/protobuf/proto"
)

func BenchmarkSint32Serialization(b *testing.B) {
	msg := &benchproto.Sint32Message{Value: []int32{-1111, -2222, -3333, -4444, -5555}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, err := proto.Marshal(msg)
		if err != nil {
			b.Fatal(err)
		}
		_ = x
	}
}

func BenchmarkInt32Serialization(b *testing.B) {
	msg := &benchproto.Int32Message{Value: []int32{-1111, -2222, -3333, -4444, -5555}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, err := proto.Marshal(msg)
		if err != nil {
			b.Fatal(err)
		}
		_ = x
	}
}

func BenchmarkSint32SerializationPositive(b *testing.B) {
	msg := &benchproto.Sint32Message{Value: []int32{1111, 2222, 3333, 4444, 5555}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, err := proto.Marshal(msg)
		if err != nil {
			b.Fatal(err)
		}
		_ = x
	}
}

func BenchmarkInt32SerializationPositive(b *testing.B) {
	msg := &benchproto.Int32Message{Value: []int32{1111, 2222, 3333, 4444, 5555}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, err := proto.Marshal(msg)
		if err != nil {
			b.Fatal(err)
		}
		_ = x
	}
}

func BenchmarkSint32SerializationMixed(b *testing.B) {
	msg := &benchproto.Sint32Message{Value: []int32{1111, -2222, 3333, -4444, 5555}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, err := proto.Marshal(msg)
		if err != nil {
			b.Fatal(err)
		}
		_ = x
	}
}

func BenchmarkInt32SerializationMixed(b *testing.B) {
	msg := &benchproto.Int32Message{Value: []int32{1111, -2222, 3333, -4444, 5555}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x, err := proto.Marshal(msg)
		if err != nil {
			b.Fatal(err)
		}
		_ = x
	}
}

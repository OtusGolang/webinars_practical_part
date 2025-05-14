// Этот тест демонстрирует, что цепочка выполняет middleware
// в нужном порядке и выдаёт ожидаемый результат.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Echo — «голый» Processor без обёрток.
func Echo(input string) string {
	return ">> " + input
}

func TestChain(t *testing.T) {
	wrapped := Chain(Echo, Trimmer, UpperCaser)

	got := wrapped("  hello  ")
	want := ">> HELLO"

	assert.Equal(t, want, got)
}

func TestChainBlock(t *testing.T) {
	wrapped := Chain(Echo, Trimmer, Blocker, UpperCaser)

	got := wrapped("  hello  ")
	want := "[blocked] hello"

	assert.Equal(t, want, got)
}

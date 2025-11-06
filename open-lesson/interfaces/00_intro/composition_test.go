package intro

import (
	"testing"
)

type Hound interface {
	destroy()
	bark(int)
}

type Retriever interface {
	Hound
	//bark(int) // duplicate method
}

func TestComposition(t *testing.T) {
	var r Retriever

	//r.bark()
	r.destroy()

	var h Hound
	h.bark(1)
	h.destroy()
}

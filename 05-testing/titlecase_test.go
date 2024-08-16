package main

import (
	"testing"

	"github.com/kulti/titlecase"
	"github.com/stretchr/testify/assert"
)

// TitleCase(str, minor) returns a str string with all words capitalized except minor words.
// The first word is always capitalized.
//
// E.g.
// TitleCase("the quick fox in the bag", "") = "The Quick Fox In The Bag"
// TitleCase("the quick fox in the bag", "in the") = "The Quick Fox in the Bag"

// Задание
// 1. Дописать существующие тесты.
// 2. Придумать один новый тест.

func TestEmpty(t *testing.T) {
	const str, minor, want = "", "", ""
	got := titlecase.TitleCase(str, minor)
	assert.Equal(t, want, got)
}

func TestWithoutMinor(t *testing.T) {
	panic("implement me")
}

func TestWithMinorInFirst(t *testing.T) {
	panic("implement me")
}

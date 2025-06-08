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
	assert.Equalf(t, want, got, "got %q, want %q", got, want)
}

func TestWithoutMinor(t *testing.T) {
	const str, minor, want = "the quick fox in the bag", "", "The Quick Fox In The Bag"
	got := titlecase.TitleCase(str, minor)
	assert.Equalf(t, want, got, "got %q, want %q", got, want)

}

func TestWithMinorInFirst(t *testing.T) {
	const str, minor, want = "the quick fox in the bag", "the", "The Quick Fox In the Bag"
	got := titlecase.TitleCase(str, minor)
	assert.Equalf(t, want, got, "got %q, want %q", got, want)
}

func TestWithSigns(t *testing.T) {
	t.Parallel()
	const str, minor, want = "the quick fox, in-the bag", "fox in the", "The Quick fox In the Bag"
	got := titlecase.TitleCase(str, minor)
	assert.Equalf(t, want, got, "got %q, want %q", got, want)
}

func TestTitleCase(t *testing.T) {
	t.Parallel()
	tt := []struct {
		name  string
		str   string
		minor string
		exp   string
	}{
		{
			name:  "empty",
			str:   "",
			minor: "",
			exp:   "",
		},
		{
			name:  "without_minor",
			str:   "the quick fox in the bag",
			minor: "",
			exp:   "The Quick Fox In The Bag",
		},
		{
			name:  "minor_in_first",
			str:   "the quick fox in the bag",
			minor: "the",
			exp:   "The Quick Fox In the Bag",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			act := titlecase.TitleCase(tc.str, tc.minor)
			assert.Equalf(t, tc.exp, act, "got %q, want %q", act, tc.exp)
		})
	}
}

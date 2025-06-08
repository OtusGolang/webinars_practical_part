package cool_stuff

import (
	"fmt"
	"strings"
	"testing"

	"github.com/kulti/titlecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) { // <-- Test...(t *testing.T)
	const s, sub, want = "chicken", "ken", 4
	got := strings.Index(s, sub)
	require.Equal(t, want, got)

	fmt.Println("hi!!!")
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

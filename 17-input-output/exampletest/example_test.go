package exampletest_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadFile(t *testing.T) {
	f, err := os.Open("testdata/example.txt")
	require.NoError(t, err)

	defer func() {
		require.NoError(t, f.Close())
	}()

	b, err := io.ReadAll(f)
	require.NoError(t, err)

	t.Log(string(b))
}

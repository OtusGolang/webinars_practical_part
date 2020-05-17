package user

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetComDomains(t *testing.T) {
	coms := getComDomains("data.dat")

	expectedBytes, err := ioutil.ReadFile("expected_coms.dat")
	require.NoError(t, err)

	var expectedComs map[string]uint32
	err = json.Unmarshal(expectedBytes, &expectedComs)
	require.NoError(t, err)
	require.Equal(t, expectedComs, coms)
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getComDomains("data.dat")
	}
}

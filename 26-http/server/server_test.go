package main

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyHandler(t *testing.T) {
	server := &MyHandler{}
	ts := httptest.NewServer(server)
	defer ts.Close()

	cases := []struct {
		name         string
		method       string
		target       string
		body         io.Reader
		responseCode int
	}{
		{"bad_request", http.MethodPost, "/vote", nil, http.StatusBadRequest},
		{
			"ok",
			http.MethodPost,
			"/vote",
			bytes.NewBufferString(`{"candidate_id": 1, "passport": "test"}`),
			http.StatusOK,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req, err := http.NewRequest(c.method, ts.URL+c.target, c.body)
			require.NoError(t, err)

			res, err := http.DefaultClient.Do(req)
			require.NoError(t, err)
			require.Equal(t, c.responseCode, res.StatusCode)
		})
	}
}

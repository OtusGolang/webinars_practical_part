package handler

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestService_SubmitVote(t *testing.T) {
	service := NewService()

	cases := []struct {
		name         string
		method       string
		target       string
		body         io.Reader
		responseCode int
	}{
		{"bad_request", http.MethodPost, "http://test.test/vote", nil, http.StatusBadRequest},
		{
			"ok",
			http.MethodPost,
			"http://test.test/vote",
			bytes.NewBufferString(`{"candidate_id": 1, "passport": "test"}`),
			http.StatusOK,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			r := httptest.NewRequest(c.method, c.target, c.body)
			w := httptest.NewRecorder()
			service.SubmitVote(w, r)
			resp := w.Result()
			require.Equal(t, c.responseCode, resp.StatusCode)
		})
	}
}

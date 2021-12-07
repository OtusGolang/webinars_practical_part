package main

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPrepareStat(t *testing.T) {
	cases := []struct {
		name      string
		code      int
		body      string
		expError  bool
		expStruct *Stat
	}{
		{"bad_request", http.StatusBadRequest, "", true, nil},
		{"ok",
			http.StatusOK,
			`{"data":{"candidate_id":1,"stat":1,"time":"2021-12-06T19:33:30.972789465+03:00"},"error":{"message":""}}`,
			false,
			&Stat{
				CandidateId: 1,
				Statistic:   1,
			}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			resRec := httptest.NewRecorder()
			_, err := resRec.WriteString(c.body)
			require.NoError(t, err)
			resRec.Code = c.code

			res, err := PrepareStat(resRec.Result())
			if c.expError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, c.expStruct.CandidateId, res.CandidateId)
				require.Equal(t, c.expStruct.Statistic, res.Statistic)
			}
		})
	}

}

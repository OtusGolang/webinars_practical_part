package golden_test

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

var update = flag.Bool("update", false, "update .golden files")

func TestGolden(t *testing.T) {
	const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Jessie
`

	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", true},
	}

	tmpl := template.Must(template.New("letter").Parse(letter))

	for _, r := range recipients {
		t.Run(r.Name, func(t *testing.T) {
			b := bytes.NewBuffer(nil)
			err := tmpl.Execute(b, r)
			require.NoError(t, err)

			goldenFileName := filepath.Join("testdata", r.Name+".golden")

			if *update {
				err := os.WriteFile(goldenFileName, b.Bytes(), 0644)
				require.NoError(t, err)
			}

			g, err := os.ReadFile(goldenFileName)
			require.NoError(t, err)

			require.Equal(t, string(g), b.String())
		})
	}
}

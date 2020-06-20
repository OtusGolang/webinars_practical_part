package golden_test

import (
	"bytes"
	"flag"
	"io/ioutil"
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
Josie
`

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	tmpl := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		t.Run(r.Name, func(t *testing.T) {
			b := bytes.NewBuffer(nil)
			err := tmpl.Execute(b, r)
			require.NoError(t, err)
			gp := filepath.Join("testdata", filepath.FromSlash(r.Name)+".golden")
			if *update {
				t.Log("update golden file")
				if err := ioutil.WriteFile(gp, b.Bytes(), 0644); err != nil {
					t.Fatalf("failed to update golden file: %s", err)
				}
			}
			g, err := ioutil.ReadFile(gp)
			if err != nil {
				t.Fatalf("failed reading .golden: %s", err)
			}
			t.Log(string(b.Bytes()))
			if !bytes.Equal(b.Bytes(), g) {
				t.Errorf("bytes do not match .golden file")
			}
		})
	}
}

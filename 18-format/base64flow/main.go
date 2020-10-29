package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

type Printer struct {
	r *io.PipeReader
	w *io.PipeWriter
}

func (p Printer) Read(d []byte) (n int, err error) {
	return p.r.Read(d)
}

func (p Printer) Write(d []byte) (n int, err error) {
	fmt.Println(string(d))
	return p.w.Write(d)
}

func main() {
	r, w := io.Pipe()
	p := Printer{r, w}

	encoder := base64.NewEncoder(base64.StdEncoding, p)
	defer func() { mustNil(encoder.Close()) }()

	go io.Copy(encoder, os.Stdin)

	decoder := base64.NewDecoder(base64.StdEncoding, p)
	io.Copy(os.Stdout, decoder)
}

func mustNil(err error) {
	if err != nil {
		panic(err)
	}
}

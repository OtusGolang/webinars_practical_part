package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	data := "Hello world"

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, err := b64.StdEncoding.DecodeString(sEnc)
	mustNil(err)
	fmt.Println(string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)

	uDec, err := b64.URLEncoding.DecodeString(uEnc)
	mustNil(err)
	fmt.Println(string(uDec))
	fmt.Println()
}

func mustNil(err error) {
	if err != nil {
		panic(err)
	}
}

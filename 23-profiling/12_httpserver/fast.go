package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	h := requestHandler
	h = fasthttp.CompressHandler(h)

	if err := fasthttp.ListenAndServe("127.0.0.1:8080", h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world!\n\n")
}

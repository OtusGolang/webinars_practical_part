package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
	//_ "net/http/pprof"
)

func main() {
	h := requestHandler
	h = fasthttp.CompressHandler(h)

	if err := fasthttp.ListenAndServe("127.0.0.1:8080", h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}

}

//func main() {
//	h := requestHandler
//	h = fasthttp.CompressHandler(h)
//
//	go func() {
//		if err := fasthttp.ListenAndServe("127.0.0.1:8080", h); err != nil {
//			log.Fatalf("Error in ListenAndServe: %s", err)
//		}
//	}()
//	mux := http.NewServeMux()
//	mux.HandleFunc("/debug/pprof/", pprof.Index)
//	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
//	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
//	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
//	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
//
//	fmt.Println("Open http://127.0.0.1:8081/debug/pprof/ in your browser")
//
//	if err := http.ListenAndServe(":8081", mux); err != nil {
//		log.Fatalf("Error in ListenAndServe: %s", err)
//	}
//
//}

func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hello, world!\n\n")
}

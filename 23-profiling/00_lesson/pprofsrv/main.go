package main

import (
	"encoding/json"
	"net/http"
	_ "net/http/pprof"
)

type Resp struct {
	Msgs []string
}

func handler(w http.ResponseWriter, r *http.Request) {
	resp := Resp{
		Msgs: make([]string, 10),
	}
	for i := 0; i < len(resp.Msgs); i++ {
		resp.Msgs[i] = "hello world"
	}
	res, _ := json.Marshal(resp)
	w.Write(res)
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":7070", nil)
}

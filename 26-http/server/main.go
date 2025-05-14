package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/OtusGolang/webinars_practical_part/26-http/handler"
)

type MyHandler struct {
	// some useful field
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serverHandler(w, r)
}

var vote = handler.NewService()

func serverHandler(w http.ResponseWriter, r *http.Request) {
	resp := &handler.Response{}
	switch r.URL.Path {
	case "/vote":
		vote.SubmitVote(w, r)
	case "/stat", "/stat/":
		vote.GetStats(w, r)
	default:
		resp.Error.Message = fmt.Sprintf("uri %s not found", r.URL.Path)
		w.WriteHeader(http.StatusNotFound)
		handler.WriteResponse(w, resp)
	}
}

// curl -d '{"candidate_id": 1, "passport": "test"}' -X POST 0.0.0.0:8080/vote
// curl 0.0.0.0:8080/stat
// curl 0.0.0.0:8080/stat/?candidate_id=1
func main() {

	handlerHttp := &MyHandler{}

	server := &http.Server{
		Addr: ":8080",
		//Handler:      http.HandlerFunc(serverHandler),
		Handler:      handlerHttp,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Print("server start on port 8080")
	log.Fatal(server.ListenAndServe())

}

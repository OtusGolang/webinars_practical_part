package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/OtusGolang/webinars_practical_part/26-http/handler"
	"github.com/OtusGolang/webinars_practical_part/26-http/middleware"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	var tr http.RoundTripper
	tr = &http.Transport{
		MaxIdleConns:    100,
		IdleConnTimeout: 90 * time.Second,
	}
	tr = middleware.NewLoggingRoundTripper(tr)

	client := http.Client{
		Transport: tr,
	}

	// post vote
	voteReq := &handler.VoteRequest{
		Passport:    "test",
		CandidateId: 1,
	}

	jsonBody, _ := json.Marshal(voteReq)
	reqVote, err := http.NewRequest(http.MethodPost, "http://0.0.0.0:8080/vote",
		bytes.NewBuffer(jsonBody))

	// respVote, err := http.DefaultClient.Do(reqVote)
	respVote, err := client.Do(reqVote)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("responce from vote: %+v\n", respVote)

	// get stat for candidate with id  1
	reqArgs := url.Values{}
	reqArgs.Add("candidate_id", "1")

	reqUrl, _ := url.Parse("http://0.0.0.0:8080/stat")
	reqUrl.RawQuery = reqArgs.Encode()

	// with context
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	reqStat, _ := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl.String(), nil)
	reqStat.Header.Add("User-Agent", `Mozilla/5.0 Gecko/20100101 Firefox/39.0`)

	// respStat, err := http.DefaultClient.Do(req)
	respStat, err := client.Do(reqStat)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("responce from stat: %+v\n", respStat)
}

type Stat struct {
	CandidateId uint32 `json:"candidate_id"`
	Statistic   int    `json:"stat"`
}

func (s *Stat) String() string {
	return ""
}

func PrepareStat(res *http.Response) (*Stat, error) {
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("state return no 200 code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response struct {
		Data *Stat `json:"data"`
	}
	
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

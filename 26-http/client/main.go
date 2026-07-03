package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/OtusGolang/webinars_practical_part/26-http/handler"
	"github.com/OtusGolang/webinars_practical_part/26-http/middleware"
	"github.com/lmittmann/tint"
)

func main() {

	//var protocols http.Protocols        // http2 enable for non-https
	//protocols.SetUnencryptedHTTP2(true) // http2 enable for non-https

	slog.SetDefault(slog.New(tint.NewHandler(os.Stdout, nil)))

	var tr http.RoundTripper
	tr = &http.Transport{
		MaxIdleConns:    100,
		IdleConnTimeout: 90 * time.Second,
		// ForceAttemptHTTP2: true,       // http2 enable for non-https
		// Protocols:         &protocols, // http2 enable for non-https
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

	jsonBody, err := json.Marshal(voteReq)
	if err != nil {
		slog.Error("error marshalling vote request", "err", err)
		return
	}

	reqVote, err := http.NewRequest(
		http.MethodPost,
		"http://0.0.0.0:8080/vote",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		slog.Error("error creating request", "err", err)
		return
	}

	// respVote, err := http.DefaultClient.Do(reqVote)
	respVote, err := client.Do(reqVote)
	if err != nil {
		slog.Error("error sending request", "err", err)
		return
	}

	slog.Info("responce from vote", "resp", respVote)

	// get stat for candidate with id 1 (path-параметр, Go 1.22+ routing)
	reqUrl := "http://0.0.0.0:8080/stat/1"

	// with context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reqStat, _ := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	reqStat.Header.Add("User-Agent", `Mozilla/5.0 Gecko/20100101 Firefox/39.0`)

	// respStat, err := http.DefaultClient.Do(req)
	respStat, err := client.Do(reqStat)
	if err != nil {
		slog.Error("error sending request", "err", err)
		return
	}

	slog.Info("responce from stat", "resp", respStat)
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

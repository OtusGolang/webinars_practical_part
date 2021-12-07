package middleware

import (
	"fmt"
	"net/http"
)

type LoggingRoundTripper struct {
	rt http.RoundTripper
}

func NewLoggingRoundTripper(rt http.RoundTripper) *LoggingRoundTripper {
	return &LoggingRoundTripper{
		rt: rt,
	}
}

func (lrt *LoggingRoundTripper) RoundTrip(req *http.Request) (res *http.Response, err error) {
	fmt.Printf("Sending request to %v\n", req.URL)

	res, err = lrt.rt.RoundTrip(req)

	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Printf("Received %v response\n", res.Status)
	}

	return
}

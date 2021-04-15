package main

import (
	"fmt"
	"net/http"
)

func main() {
	type Result struct {
		Error    error
		Response *http.Response
	}

	checkStatus := func(done <-chan struct{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result

				result.Response, result.Error = http.Get(url)

				//resp, err := http.Get(url)
				//
				//if err != nil {
				//	result = Result{Error: err, Response: nil}
				//} else {
				//	result = Result{Error: nil, Response: resp}
				//}

				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan struct{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response.Status)
	}
}

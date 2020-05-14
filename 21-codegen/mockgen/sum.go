package mockgen

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetPage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to GET %q: %w", url, err)
	}
	result, err := ioutil.ReadAll(resp.Body)
	return result, nil
}

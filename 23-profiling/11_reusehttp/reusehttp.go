package reusehttp

import (
	"io/ioutil"
	"net/http"
	"time"
)

var (
	client = http.Client{
		Timeout:   3 * time.Second,
		Transport: &http.Transport{},
	}
)

func Slow() ([]byte, error) {
	resp, err := http.Get("https://api.vk.com/method/users.get")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, nil
}

func Fast() ([]byte, error) {
	resp, err := client.Get("https://api.vk.com/method/users.get")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, nil
}

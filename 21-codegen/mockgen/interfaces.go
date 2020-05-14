package mockgen

import "net/http"

//go:generate mockgen -source=$GOFILE -destination ./mocks/mock_getter.go -package mocks Getter
type Getter interface {
	Get(url string) (resp *http.Response, err error)
}

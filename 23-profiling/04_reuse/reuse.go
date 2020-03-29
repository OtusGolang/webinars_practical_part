package reuse

import "encoding/json"

type A struct {
	I int
}

func Slow() {
	for i := 0; i < 1000; i++ {
		a := &A{}
		json.Unmarshal([]byte("{\"i\": 32}"), a)
	}
}

func Fast() {
	a := &A{}
	for i := 0; i < 1000; i++ {
		json.Unmarshal([]byte("{\"i\": 32}"), a)
	}
}
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	j := []byte(`{"Name":"Vasya",
                  "Job":{"Department":"Operations","Title":"Boss"}}`)

	var p2 interface{}
	json.Unmarshal(j, &p2)
	fmt.Printf("p2: %v\n", p2)

	person := p2.(map[string]interface{})
	fmt.Printf("name=%s\n", person["Name"])
}

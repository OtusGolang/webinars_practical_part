package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) MarshalJSON() ([]byte, error) {
	fmt.Println(p.Name + " is cool in marshaling")
	return json.Marshal(p)
}

func (p *Person) UnmarshalJSON(data []byte) error {
	fmt.Println(p.Name + " is cool in unmarshaling")
	return json.Unmarshal(data, p)
}

func main() {
	var p Person
	if err := json.Unmarshal([]byte(`{"Name": "Vasya", "Age": 32}`), &p); err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}

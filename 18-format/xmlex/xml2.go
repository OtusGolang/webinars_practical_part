package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Email struct {
	Where string `xml:"where,attr"`
	Addr  string
}

type Address struct {
	City, State string
}

type Result struct {
	XMLName xml.Name `xml:"Person"`
	Name    string   `xml:"FullName"`
	Phone   string
	Email   []Email
	Groups  []string `xml:"Group>Value"`
	Address
}

func main() {
	v := Result{Name: "none", Phone: "none"}

	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`

	if err := xml.Unmarshal([]byte(data), &v); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", v)
}

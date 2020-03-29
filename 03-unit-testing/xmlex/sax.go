package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func main() {

	data := `
		<Person id="34">
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
		<Person id="35">
			<FullName>Rob Pike</FullName>
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
		<Person id="35">
			<FullName>Russ Cox</FullName>
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

	decoder := xml.NewDecoder(strings.NewReader(data))

	inFullName := false

	names := []string{}

	for {
		token, err := decoder.Token()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}

		switch se := token.(type) {
		case xml.Attr:
			fmt.Printf("Attribute: %v Value %s\n", se.Name.Local, se.Value)
		case xml.StartElement:
			fmt.Printf("Start element: %v Attr %s\n", se.Name.Local, se.Attr)
			inFullName = se.Name.Local == "FullName"
		case xml.EndElement:
			fmt.Printf("End element: %v\n", se.Name.Local)
			inFullName = false
		case xml.Name:
			fmt.Printf("Name element: %v\n", se.Local)
		case xml.CharData:
			fmt.Printf("Data element: %v\n", string(se))
			if inFullName {
				names = append(names, string(se))
			}
		default:
			fmt.Printf("Unhandled element: %T", se)
		}

	}

	fmt.Printf("All names: %v", names)

}

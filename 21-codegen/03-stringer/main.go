package main

import "fmt"

//go:generate stringer -type=MessageStatus
type MessageStatus int

const (
	Sent MessageStatus = iota
	Received
	Rejected
)

func main() {
	status := Sent
	fmt.Printf("Message is %s\n", status) // Message is Sent
}

// run with: go run main.go messagestatus_string.go

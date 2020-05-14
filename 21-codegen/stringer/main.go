package main

import "fmt"

//go:generate stringer -type=MessageStatus
type MessageStatus int

const (
	Sent MessageStatus = iota
	Received
	Rejected
)

var _ fmt.Stringer = (*MessageStatus)(nil)

func main() {
	status := Sent
	fmt.Printf("Message is %s", status) // Message is Sent
}

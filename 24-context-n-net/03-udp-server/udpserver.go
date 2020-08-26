package main

import (
	"fmt"
	"log"
	"net"
)

//TODO: Make worker

//TODO: SHOW BROADCAST

func main() {
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:3303")

	l, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Cannot listen")
	}

	defer l.Close()

	msg := make([]byte, 1024)
	for {
		length, fromAddr, err := l.ReadFromUDP(msg)
		if err != nil {
			log.Fatalf("Error happened")
		}

		fmt.Printf("Message from %s with length %d: %s\n", fromAddr.String(), length, string(msg))
	}
}

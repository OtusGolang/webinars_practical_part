package main

import (
	"log"
	"net"
	"time"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:3303")
	if err != nil {
		log.Fatalf("cannot resolve server addr: %v", err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		log.Fatalf("cannot dial to server: %v", err)
	}

	order := 0
	for {
		msg := ""
		switch order {
		case 0:
			msg = "qwertyuiop"
			order++
		case 1:
			msg = "asdfghjkl"
			order++
		case 2:
			msg = "zxcvbnm"
			order = 0
		}

		_, err := conn.Write([]byte(msg))
		if err != nil {
			log.Fatalf("cannot send: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn)  {
	defer conn.Close()
	conn.Write([]byte(fmt.Sprintf("Welcome to %s, friend from %s\n", conn.LocalAddr(), conn.RemoteAddr())))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		log.Printf("RECEIVED: %s", text)
		if text == "quit" || text == "exit" {
			break
		}

		conn.Write([]byte(fmt.Sprintf("I have received '%s'\n", text)))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error happend on connection with %s: %v", conn.RemoteAddr(), err)
	}

	log.Printf("Closing connection with %s", conn.RemoteAddr())

}

//TODO multithread

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:3302")
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("Cannot accept: %v", err)
		}

		handleConnection(conn)
	}
}

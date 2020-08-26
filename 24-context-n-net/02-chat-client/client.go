package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func readRoutine(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(conn)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				log.Printf("CANNOT SCAN")
				break OUTER
			}
			text := scanner.Text()
			log.Printf("From server: %s", text)
		}
	}
	log.Printf("Finished readRoutine")
}

func writeRoutine(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		default:
			if !scanner.Scan() {
				break OUTER
			}
			str := scanner.Text()
			log.Printf("To server %v\n", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}

	}
	log.Printf("Finished writeRoutine")
}

//TODO:

func main() {
	dialer := &net.Dialer{}
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)

	conn, err := dialer.DialContext(ctx, "tcp", "127.0.0.1:3302")
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	go readRoutine(ctx, conn)
	go writeRoutine(ctx, conn)

	time.Sleep(1 * time.Minute)
	cancel()
	time.Sleep(1 * time.Minute)
	conn.Close()
}

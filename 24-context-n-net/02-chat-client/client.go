package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func readRoutine(ctx context.Context, conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
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

func writeRoutine(ctx context.Context, conn net.Conn, wg *sync.WaitGroup, stdin chan string) {
	defer wg.Done()
	//scanner := bufio.NewScanner(os.Stdin)
OUTER:
	for {
		select {
		case <-ctx.Done():
			break OUTER
		case str := <- stdin:
			//if !scanner.Scan() {
			//	break OUTER
			//}
			//str := scanner.Text()
			log.Printf("To server %v\n", str)

			conn.Write([]byte(fmt.Sprintf("%s\n", str)))
		}

	}
	log.Printf("Finished writeRoutine")
}

func stdoutScan() chan string {
	out := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			out<-scanner.Text()
		}
		if scanner.Err() != nil {
			close(out)
		}
	}()
	return out
}

//TODO:

func main() {
	dialer := &net.Dialer{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

	conn, err := dialer.DialContext(ctx, "tcp", "127.0.0.1:3302")
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	stdin := stdoutScan()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(){
		readRoutine(ctx, conn, wg)
		cancel()
	}()

	wg.Add(1)
	go func(){
		writeRoutine(ctx, conn, wg, stdin)
	}()

	wg.Wait()
	conn.Close()
}

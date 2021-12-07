package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func readLoop(c *websocket.Conn) {
	for {
		_, buff, err := c.ReadMessage()
		if err != nil {
			c.Close()
			break
		}
		fmt.Printf("%s\n", buff)
	}
}

func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM)

	url := "ws://localhost:8080/stat-stream"
	wsDialer := websocket.Dialer{}
	conn, _, err := wsDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}

	go readLoop(conn)
	<-ch
	conn.Close()
}

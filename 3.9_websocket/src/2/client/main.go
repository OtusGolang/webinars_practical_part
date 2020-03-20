package main

import (
	"bufio"
	"log"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func runWriteLoop(conn *websocket.Conn, ch <-chan []byte) {
	for msg := range ch {
		conn.SetWriteDeadline(time.Now().Add(5 * time.Second))
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			log.Fatal(err)
		}
		log.Printf("message was sent")
	}
}

func main() {
	url := "ws://localhost:33333/ws?cid=101"
	wsDialer := websocket.Dialer{}
	conn, _, err := wsDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan []byte)
	go runWriteLoop(conn, ch)

	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		ch <- []byte(text)
	}
}

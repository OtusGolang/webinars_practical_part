package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/websocket"
	"github.com/lmittmann/tint"
)

func readLoop(c *websocket.Conn) {
	for {
		_, buff, err := c.ReadMessage()
		if err != nil {
			c.Close()
			break
		}
		slog.Info("got datagram", "content", string(buff))
	}
}

func main() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stdout, nil)))

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM)

	url := "ws://localhost:8080/stat-stream"
	wsDialer := websocket.Dialer{}
	conn, _, err := wsDialer.Dial(url, nil)
	if err != nil {
		slog.Error("error dialing", "err", err)
		return
	}

	go readLoop(conn)
	<-ch
	conn.Close()
}

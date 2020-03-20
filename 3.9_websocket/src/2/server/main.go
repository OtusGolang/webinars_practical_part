package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	clientsMutex sync.Mutex
	clients      map[string]*websocket.Conn = make(map[string]*websocket.Conn)

	upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
)

func readLoop(c *websocket.Conn, cid string) {
	for {
		_, buff, err := c.ReadMessage()
		if err != nil {
			c.Close()
			break
		}

		log.Printf("received message: cid=%s, message=%s", cid, string(buff))
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	cid := r.URL.Query().Get("cid")
	if cid == "" {
		http.Error(w, "cid not passed", 400)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	conn.SetCloseHandler(func(int, string) error {
		log.Printf("remove connection, cid=%s", cid)
		clientsMutex.Lock()
		delete(clients, cid)
		clientsMutex.Unlock()
		return nil
	})

	log.Printf("add connection, cid=%s", cid)
	clientsMutex.Lock()
	clients[cid] = conn
	clientsMutex.Unlock()

	go readLoop(conn, cid)
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	log.Fatal(http.ListenAndServe("localhost:33333", nil))
}

package main

import (
	"io/ioutil"
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

func readLoop(c *websocket.Conn) {
	for {
		if _, _, err := c.NextReader(); err != nil {
			c.Close()
			break
		}
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

	go readLoop(conn)
}

func pushHandler(w http.ResponseWriter, r *http.Request) {
	cid := r.URL.Query().Get("cid")
	if cid == "" {
		http.Error(w, "cid not passed", 400)
		return
	}

	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read message", 400)
		return
	}

	clientsMutex.Lock()
	conn, ok := clients[cid]
	clientsMutex.Unlock()
	if !ok {
		http.NotFound(w, r)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/push", pushHandler)

	log.Fatal(http.ListenAndServe("localhost:33333", nil))
}

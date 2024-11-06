package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader    = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	clients     = make(map[*websocket.Conn]bool)
	broadcast   = make(chan float64, 100)
	clientMutex sync.Mutex
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
			log.Println("Error upgrading connection:", err)
			return
	}
	
	defer func() {
			clientMutex.Lock()
			delete(clients, conn)
			clientMutex.Unlock()
			conn.Close()
	}()

	clientMutex.Lock()
	clients[conn] = true
	clientMutex.Unlock()

	for {
			if _, _, err := conn.ReadMessage(); err != nil {
					log.Println("Error reading from client:", err)
					break
			}
	}
}

func handleBroadcast() {
	for avgPrice := range broadcast {
		clientMutex.Lock()
		for client := range clients {
			err := client.WriteJSON(map[string]float64{"averagePrice": avgPrice})
			if err != nil {
				log.Println("Error writing to client:", err)
				client.Close()
				delete(clients, client)
			}
		}
		clientMutex.Unlock()
	}
}

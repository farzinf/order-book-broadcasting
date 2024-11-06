package main

import (
	"log"
	"net/http"
)

func main() {	
	http.HandleFunc("/ws", handleConnections)
	go handleBroadcast()

	go connectToBinanceAndProcess()

	log.Println("WebSocket server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow connections from any origin
		return true
	},
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	// Handle the WebSocket connection
	// Implement your WebSocket logic here
	// You can read and write messages using the 'conn' object

	// Example: Echo server
	for {
		// Read message from the client
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}

		// Echo the message back to the client
		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Failed to write message:", err)
			break
		}
	}

	// Close the WebSocket connection
	conn.Close()
}

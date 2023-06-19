package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader    = websocket.Upgrader{}
	connections = make(map[string]*websocket.Conn)
	mutex       = sync.Mutex{}
)

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	// Get the user's nickname from the request
	nickname, _ := nicknameFromSession(r)

	// Add the WebSocket connection to the connections map
	mutex.Lock()
	connections[nickname] = conn
	mutex.Unlock()

	for {
		// Read message from the client
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}
		// Extract the recipient user's nickname from the message
		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("Failed to unmarshal message:", err)
			break
		}
		// Handle the received message
		handleMessage(msg.NicknameTo, string(message))
	}

	// Remove the WebSocket connection from the connections map when the connection is closed
	mutex.Lock()
	delete(connections, nickname)
	mutex.Unlock()
}

func handleMessage(nickname, message string) {
	// Check if the recipient user has an active WebSocket connection
	mutex.Lock()
	conn, found := connections[nickname]
	mutex.Unlock()

	if found {
		// Send the message to the recipient user's WebSocket connection
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Println("Failed to write message:", err)
		}
	} else {
		log.Println("No active WebSocket connection found for", nickname)
	}
}

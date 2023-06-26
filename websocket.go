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
	// in order to enable full-duplex communication and support WebSocket-specific features, the HTTP connection needs to be upgraded to a WebSocket connection.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	// Get the user's nickname from the request
	nickname, _ := nicknameFromSession(r)

	// Add the WebSocket connection to the connections map
	//map is used to maintain active WebSocket connections.
	mutex.Lock()
	connections[nickname] = conn
	mutex.Unlock()

	//The function enters a loop to continuously read messages from the client.
	for {
		// Read message from the client
		_, message, err := conn.ReadMessage()
		//If an error occurs during reading, it logs the error and breaks out of the loop.
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}
		//It then unmarshals the received message into a Message struct, which includes the recipient user's nickname.
		//If unmarshaling fails, it logs the error and breaks out of the loop.
		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("Failed to unmarshal message:", err)
			break
		}
		//Finally, it calls the handleMessage function, passing the recipient user's nickname and the message as parameters to handle the received message.
		handleMessage(r, msg.NicknameTo, msg.NicknameFrom, msg)
	}

	// Remove the WebSocket connection from the connections map when the connection is closed
	//The function uses a mutex to protect concurrent access to the connections map to ensure thread safety.
	mutex.Lock()
	delete(connections, nickname)
	mutex.Unlock()
}

func handleMessage(r *http.Request, nickname string, senderNickname string, message Message) {
	// Check if the recipient user has an active WebSocket connection
	mutex.Lock()
	recipientConn, recipientFound := connections[nickname]
	mutex.Unlock()

	// Check if the sender user has an active WebSocket connection
	senderNickname, _ = nicknameFromSession(r)
	mutex.Lock()
	senderConn, senderFound := connections[senderNickname]
	mutex.Unlock()

	if recipientFound {
		// Send the message to the recipient user's WebSocket connection
		data, err := json.Marshal(message)
		if err != nil {
			log.Println("Failed to marshal message:", err)
			return
		}
		err = recipientConn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Failed to write message to recipient:", err)
		}
	} else {
		log.Println("No active WebSocket connection found for recipient:", nickname)
	}

	if senderFound {
		// Send the message to the sender's WebSocket connection
		data, err := json.Marshal(message)
		if err != nil {
			log.Println("Failed to marshal message:", err)
			return
		}
		err = senderConn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println("Failed to write message to sender:", err)
		}
	} else {
		log.Println("No active WebSocket connection found for sender:", senderNickname)
	}
}

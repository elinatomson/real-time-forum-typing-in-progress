package handlers

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

func Websocket(w http.ResponseWriter, r *http.Request) {
	//upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}

	nickname, err := nicknameFromSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//add the WebSocket connection to the connections map to maintain active WebSocket connections.
	mutex.Lock()
	connections[nickname] = conn
	mutex.Unlock()

	//the function enters a loop to continuously read messages from the client.
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message:", err)
			break
		}
		//unmarshals the received message into a Message struct, which includes the recipient user's nickname.
		var msg Message
		err = json.Unmarshal(message, &msg)
		if err != nil {
			log.Println("Failed to unmarshal message:", err)
			break
		}
		//finally calling the handleMessage function, passing the recipient user's nickname and the message as parameters to handle the received message
		handleMessage(r, w, msg.NicknameTo, msg)
	}

	//remove the WebSocket connection from the connections map when the connection is closed
	mutex.Lock()
	delete(connections, nickname)
	mutex.Unlock()
}

func handleMessage(r *http.Request, w http.ResponseWriter, nickname string, message Message) {
	senderNickname, err := nicknameFromSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//checking if both users has an active WebSocket connection
	mutex.Lock()
	recipientConn, recipientFound := connections[nickname]
	senderConn, senderFound := connections[senderNickname]
	mutex.Unlock()

	if recipientFound {
		//if the sender is typing (message.Typing is true), then it is needed to display the typing message for the recipient
		if message.Typing {
			typingMessage := Message{
				Typing:       true,
				NicknameTo:   nickname,
				NicknameFrom: senderNickname,
			}
			data, err := json.Marshal(typingMessage)
			if err != nil {
				log.Println("Failed to marshal typing message:", err)
				return
			}
			//sending typing message to the recipient user's WebSocket connection
			err = recipientConn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println("Failed to write typing message to recipient:", err)
			}
			//and also sending the message itself to the recipient user's WebSocket connection
		} else {
			data, err := json.Marshal(message)
			if err != nil {
				log.Println("Failed to marshal message:", err)
				return
			}
			err = recipientConn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println("Failed to write message to recipient:", err)
			}
		}
	} else {
		log.Println("No active WebSocket connection found for recipient:", nickname)
	}

	if senderFound {
		//sending the message to the sender's WebSocket connection
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

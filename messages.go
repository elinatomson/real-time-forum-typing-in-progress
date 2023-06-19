package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Message struct {
	MessageID    int
	Message      string    `json:"message"`
	NicknameFrom string    `json:"nicknamefrom"`
	NicknameTo   string    `json:"nicknameto"`
	Date         time.Time `json:"date"`
}

type IncomingMessage struct {
	Message      string `json:"message"`
	NicknameFrom string `json:"nicknamefrom"`
	NicknameTo   string `json:"nicknameto"`
}

var message Message

func messageing(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		addMessage(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Error 405, method not allowed")
		return
	}
}

func addMessage(w http.ResponseWriter, r *http.Request) {
	var incomingMessage IncomingMessage

	err := json.NewDecoder(r.Body).Decode(&incomingMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := Message{
		Message:      incomingMessage.Message,
		NicknameFrom: incomingMessage.NicknameFrom,
		NicknameTo:   incomingMessage.NicknameTo,
		Date:         time.Now(),
	}

	message.NicknameFrom, _ = nicknameFromSession(r)

	if message.Message != "" {
		// Insert the message for the sender
		_, err := db.Exec(`INSERT INTO messages (message, nicknamefrom, nicknameto, date) VALUES (?, ?, ?, ?)`, message.Message, message.NicknameFrom, message.NicknameTo, message.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Convert the message to a JSON string
		jsonMsg, err := json.Marshal(message)
		if err != nil {
			log.Println("Failed to marshal message:", err)
			return
		}

		// Send the message to the recipient user's WebSocket connection
		handleMessage(message.NicknameTo, string(jsonMsg))
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Message sent successfully")
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	nicknameTo := r.URL.Query().Get("nicknameTo")
	nicknameFrom, _ := nicknameFromSession(r)

	rows, err := db.Query(`
		SELECT message, nicknamefrom, nicknameto, date FROM messages WHERE (nicknameto = ? AND nicknamefrom = ?) OR (nicknameto = ? AND nicknamefrom = ?)`, nicknameTo, nicknameFrom, nicknameFrom, nicknameTo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.Message, &msg.NicknameFrom, &msg.NicknameTo, &msg.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messages = append(messages, msg)
	}

	// Send the messages as a JSON response
	json.NewEncoder(w).Encode(messages)
}

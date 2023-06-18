package main

import (
	"encoding/json"
	"fmt"
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
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message.Date = time.Now()
	message.NicknameFrom, _ = nicknameFromSession(r)

	if message.Message != "" {
		_, err := db.Exec(`INSERT INTO messages (message, nicknamefrom, nicknameto, date) VALUES (?, ?, ?, ?)`, message.Message, message.NicknameFrom, message.NicknameTo, message.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Message sent successfully")
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	nicknameTo := r.URL.Query().Get("nicknameTo")
	nicknameFrom, _ := nicknameFromSession(r)

	rows, err := db.Query(`SELECT message, nicknamefrom  FROM messages WHERE nicknameto = ? AND nicknamefrom = ?`, nicknameTo, nicknameFrom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []Message
	var message Message
	for rows.Next() {
		err := rows.Scan(&message.Message, &message.NicknameFrom)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messages = append(messages, message)
	}
	//send the posts as a JSON response
	json.NewEncoder(w).Encode(messages)
}

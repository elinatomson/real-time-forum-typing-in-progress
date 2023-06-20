package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Message sent successfully")
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	pageSize := r.URL.Query().Get("pageSize")
	//convert page and pageSize to integers
	pageNum, _ := strconv.Atoi(page)
	pageSizeNum, _ := strconv.Atoi(pageSize)
	//calculate the offset based on the page number and page size
	offset := (pageNum - 1) * pageSizeNum

	nicknameTo := r.URL.Query().Get("nicknameTo")
	nicknameFrom, _ := nicknameFromSession(r)

	rows, err := db.Query(`
		SELECT message, nicknamefrom, nicknameto, date FROM messages WHERE (nicknameto = ? AND nicknamefrom = ?) OR (nicknameto = ? AND nicknamefrom = ?)
		ORDER BY date DESC LIMIT ? OFFSET ?`, nicknameTo, nicknameFrom, nicknameFrom, nicknameTo, pageSizeNum, offset)
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

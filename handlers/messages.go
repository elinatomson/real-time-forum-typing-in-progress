package handlers

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

var message Message

func Messageing(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		addMessage(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Error 405, method not allowed")
		return
	}
}

func addMessage(w http.ResponseWriter, r *http.Request) {

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message = Message{
		Message:      message.Message,
		NicknameFrom: message.NicknameFrom,
		NicknameTo:   message.NicknameTo,
		Date:         time.Now(),
	}

	message.NicknameFrom, _ = nicknameFromSession(r)

	if message.Message != "" {
		//insert the message
		_, err := db.Exec(`INSERT INTO messages (message, nicknamefrom, nicknameto, date) VALUES (?, ?, ?, ?)`, message.Message, message.NicknameFrom, message.NicknameTo, message.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//update the last_message_date
		_, err = db.Exec("UPDATE users SET last_message_date = DATETIME('now') WHERE nickname = ?", message.NicknameFrom)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	fmt.Printf(message.NicknameFrom)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Message sent successfully")
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
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

func UnreadMessages(w http.ResponseWriter, r *http.Request) {
	nicknameTo, _ := nicknameFromSession(r)
	rows, err := db.Query("SELECT message, nicknamefrom, nicknameto, date FROM messages WHERE read = 0 AND nicknameto = ?", nicknameTo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	//iterate over the query results and build the list of unread messages
	var unreadMessages = []Message{}
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.Message, &msg.NicknameFrom, &msg.NicknameTo, &msg.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		unreadMessages = append(unreadMessages, msg)
	}
	json.NewEncoder(w).Encode(unreadMessages)
}

func MessagesAsRead(w http.ResponseWriter, r *http.Request) {
	nicknameTo, _ := nicknameFromSession(r)
	nicknameFrom := r.URL.Query().Get("nicknameFrom")
	_, err := db.Exec(`UPDATE messages SET read = 1 WHERE nicknameto = ? AND nicknamefrom = ? AND read = 0`, nicknameTo, nicknameFrom)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "All messages sent for the logged in user marked as read.")
}

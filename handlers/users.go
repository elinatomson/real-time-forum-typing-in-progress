package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	//users must be organized by the last message sent. If the user is new and does not present messages you must organize it in alphabetic order.
	//so taking into account the last_message_date from users table and sorting is in mesages.js file in usersForChat function
	//also comparing with excisting sessions to show the user online or offline
	rows, err := db.Query(`
		SELECT users.nickname, (CASE WHEN sessions.nickname IS NULL THEN FALSE ELSE TRUE END) AS online, users.last_message_date
		FROM users
		LEFT JOIN sessions ON users.nickname = sessions.nickname
		LEFT JOIN (
			SELECT nicknamefrom, MAX(date) AS last_message_date
			FROM messages
			GROUP BY nicknamefrom
		) AS last_message ON users.nickname = last_message.nicknamefrom
		ORDER BY COALESCE(last_message.last_message_date, '1900-01-01') DESC, users.nickname ASC
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		var user User
		var lastMessageDate sql.NullTime
		err := rows.Scan(&user.Nickname, &user.Online, &lastMessageDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if lastMessageDate.Valid {
			user.LastMessageDate = lastMessageDate.Time
		}

		users = append(users, user)
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

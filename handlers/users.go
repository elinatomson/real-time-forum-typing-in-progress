package handlers

import (
	"encoding/json"
	"net/http"
	"time"
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
		panic(err)
	}

	users := make([]User, 0)

	for rows.Next() {
		var nickname string
		var online bool
		var lastMessageDate time.Time
		rows.Scan(&nickname, &online, &lastMessageDate)

		user := User{
			Nickname:        nickname,
			Online:          online,
			LastMessageDate: lastMessageDate,
		}
		users = append(users, user)
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

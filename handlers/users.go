package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sort"
	"time"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	//users must be organized by the last message sent. If the user is new and does not present messages you must organize it in alphabetic order.
	//so taking into account the last_message_date from users table and sorting is in messages.js file in usersForChat function
	//also comparing with excisting sessions to show the user online or offline
	currentUser, err := nicknameFromSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows, err := db.Query(`
	SELECT users.nickname, (CASE WHEN sessions.nickname IS NULL THEN FALSE ELSE TRUE END) AS online, MAX(messages.date) AS last_message_date
	FROM users
	LEFT JOIN sessions ON users.nickname = sessions.nickname
	LEFT JOIN messages ON (users.nickname = messages.nicknamefrom OR users.nickname = messages.nicknameto)
		AND (messages.nicknamefrom = ? OR messages.nicknameto = ?)
	GROUP BY users.nickname
	ORDER BY COALESCE(last_message_date, '1900-01-01') DESC, users.nickname ASC
`, currentUser, currentUser, currentUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		var user User
		var lastMessageDate sql.NullString
		err := rows.Scan(&user.Nickname, &user.Online, &lastMessageDate)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if lastMessageDate.Valid {
			layout := "2006-01-02 15:04:05.999999999-07:00"
			user.LastMessageDate, err = time.Parse(layout, lastMessageDate.String)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		users = append(users, user)
	}

	SortUsers(users)

	//setting the currentUser in the users db table as a true to add the current user's nickname to the response
	for i := range users {
		if users[i].Nickname == currentUser {
			users[i].CurrentUser = true
			break
		}
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

func SortUsers(users []User) {
	sort.Slice(users, func(i, j int) bool {
		if users[i].LastMessageDate.IsZero() && users[j].LastMessageDate.IsZero() {
			return users[i].Nickname < users[j].Nickname
		}
		if users[i].LastMessageDate.IsZero() {
			return false
		}
		if users[j].LastMessageDate.IsZero() {
			return true
		}
		if users[i].LastMessageDate.Equal(users[j].LastMessageDate) {
			return users[i].Nickname < users[j].Nickname
		}
		return users[i].LastMessageDate.After(users[j].LastMessageDate)
	})
}

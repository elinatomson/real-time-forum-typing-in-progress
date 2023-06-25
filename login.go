package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Nickname        string    `json:"nickname"`
	Password        string    `json:"password"`
	Online          bool      `json:"online"`
	LastMessageDate time.Time `json:"last_message_date"`
}

func logIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		logInUser(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Error 405, method not allowed")
		return
	}
}

func logInUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = checkUser(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func checkUser(w http.ResponseWriter, user User) error {
	if sessionExists(db, user.Nickname) {
		_, err := db.Exec(`DELETE FROM sessions WHERE nickname = ?`, user.Nickname)
		if err != nil {
			return err
		}
	}
	stmt := `SELECT password FROM users WHERE nickname = ?`
	row := db.QueryRow(stmt, user.Nickname)
	var hash string
	err := row.Scan(&hash)
	if user.Nickname == "" || user.Password == "" {
		return errors.New("Please insert nickname and password!")
	}
	if err != nil {
		return errors.New("Nickname or password is not correct!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password))
	if err == nil {
		addCookie(w, user.Nickname)
		return nil
	}
	return errors.New("Nickname or password is not correct!")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
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

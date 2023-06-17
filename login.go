package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Online   bool   `json:"online"`
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
	//query the users table and join with the sessions table to check online status, meaning if the same nickname has a session, then he is as online.
	rows, err := db.Query(`
		SELECT users.nickname,(CASE WHEN sessions.nickname IS NULL THEN FALSE ELSE TRUE END) AS online
		FROM users LEFT JOIN sessions ON users.nickname = sessions.nickname
	`)
	if err != nil {
		panic(err)
	}

	users := make([]User, 0)

	//iterate over the query results and build the user list
	for rows.Next() {
		var nickname string
		var online bool
		rows.Scan(&nickname, &online)

		user := User{
			Nickname: nickname,
			Online:   online,
		}
		users = append(users, user)
	}

	//convert the user list to JSON
	jsonData, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	//set the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//write the JSON data to the response
	w.Write(jsonData)
}

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
		_, err := db.Exec(`DELETE FROM sessions WHERE username = ?`, user.Nickname)
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

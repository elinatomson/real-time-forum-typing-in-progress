package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"real-time-forum/backend/database"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Nickname        string    `json:"nickname"`
	Password        string    `json:"password"`
	Online          bool      `json:"online"`
	LastMessageDate time.Time `json:"last_message_date"`
	CurrentUser     bool      `json:"currentUser"`
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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
	if user.Nickname == "" || user.Password == "" {
		return errors.New("Please insert nickname and password")
	}

	if sessionExists(database.Db, user.Nickname) {
		_, err := database.Db.Exec(`DELETE FROM sessions WHERE nickname = ?`, user.Nickname)
		if err != nil {
			return err
		}
	}

	stmt := `SELECT password FROM users WHERE nickname = ?`
	row := database.Db.QueryRow(stmt, user.Nickname)
	var hash string
	err := row.Scan(&hash)
	if err != nil {
		return errors.New("Nickname or password is not correct!")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(user.Password))
	if err != nil {
		return errors.New("Nickname or password is not correct")
	}
	addCookie(w, user.Nickname)
	return nil
}

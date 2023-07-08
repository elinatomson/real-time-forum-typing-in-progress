package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"real-time-forum/backend/database"

	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	Nickname  string `json:"nickname"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		addUserData(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Error 405, method not allowed")
		return
	}
}

func addUserData(w http.ResponseWriter, r *http.Request) {
	var data UserData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = insertUserData(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func insertUserData(userData UserData) error {
	//checking if nickname or email already exists
	stmt := `SELECT email FROM users WHERE email = ?`
	row := database.Db.QueryRow(stmt, userData.Email)
	var email string
	err := row.Scan(&email)
	if err != sql.ErrNoRows {
		return errors.New("Email already taken")
	}

	stmt = `SELECT nickname FROM users WHERE nickname = ?`
	row = database.Db.QueryRow(stmt, userData.Nickname)
	var nickname string
	err = row.Scan(&nickname)
	if err != sql.ErrNoRows {
		return errors.New("Nickname already taken")
	}

	//hashing password
	hash, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	//inserting data do users table
	_, err = database.Db.Exec("INSERT INTO users(nickname, age, gender, firstname, lastname, email, password) VALUES(?, ?, ?, ?, ?, ?, ?)",
		userData.Nickname, userData.Age, userData.Gender, userData.FirstName, userData.LastName, userData.Email, hash)
	if err != nil {
		return err
	}
	return nil
}

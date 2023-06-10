package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

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

var userData []UserData

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
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
	row := db.QueryRow(stmt, userData.Email)
	err := row.Scan(&userData.Email)
	if err != sql.ErrNoRows {
		return errors.New("Email already taken")
	}
	stmt1 := `SELECT nickname FROM users WHERE nickname = ?`
	row1 := db.QueryRow(stmt1, userData.Nickname)
	err1 := row1.Scan(&userData.Nickname)
	if err1 != sql.ErrNoRows {
		return errors.New("Nickname already taken")
	}

	//hashing password
	var hash []byte
	hash, _ = bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)

	//inserting data do users table
	_, err3 := db.Exec("INSERT INTO users(nickname, age, gender, firstname, lastname, email, password) VALUES(?, ?, ?, ?, ?, ?, ?)", userData.Nickname, userData.Age, userData.Gender, userData.FirstName, userData.LastName, userData.Email, hash)
	if err3 != nil {
		return err3
	}
	return nil
}

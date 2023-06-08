package main

import (
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
	var count int
	err1 := db.QueryRow("SELECT COUNT(*) FROM users WHERE nickname = ? ", userData.Nickname).Scan(&count)
	if err1 != nil {
		return err1
	}

	if count > 0 {
		return errors.New("Nickname already taken")
	}

	err2 := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", userData.Email).Scan(&count)
	if err2 != nil {
		return err2
	}

	if count > 0 {
		return errors.New("Email already taken")
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

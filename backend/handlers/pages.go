package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("Page not found")
		http.Error(w, "This page does not exist.", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("frontend/index.html")
	if err != nil {
		log.Println("Failed to parse template:", err)
		http.Error(w, "Something went wrong, please try again.", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Failed to execute template:", err)
		http.Error(w, "Something went wrong, please try again.", http.StatusInternalServerError)
		return
	}
}

func UserPage(w http.ResponseWriter, r *http.Request) {
	nickname, err := nicknameFromSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := User{
		Nickname: nickname,
	}

	jsonResponse, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

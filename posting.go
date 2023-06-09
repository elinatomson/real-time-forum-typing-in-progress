package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Post struct {
	ID           int
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Movies       string    `json:"movies"`
	Serials      string    `json:"serials"`
	Realityshows string    `json:"realityshows"`
	Date         time.Time `json:"date"`
	Nickname     string    `json:"nickname"`
	Comments     []Comment
}

var posts []Post

func posting(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		addPost(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Error 405, method not allowed")
		return
	}
}

func addPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var nickname, _ = nicknameFromSession(r)

	if post.Title != "" && post.Content != "" {
		_, err := db.Exec(`INSERT INTO posts (title, content, category1, category2, category3, date, nickname) VALUES (?, ?, ?, ?, ?, ?, ?)`, post.Title, post.Content, post.Movies, post.Serials, post.Realityshows, post.Date, nickname)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	//back to the homepage after successful post creation
	http.Redirect(w, r, "/", http.StatusFound)
}

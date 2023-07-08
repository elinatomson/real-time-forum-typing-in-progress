package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"real-time-forum/backend/database"
	"time"
)

type Post struct {
	ID           int
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Movies       string    `json:"movies"`
	Serials      string    `json:"serials"`
	RealityShows string    `json:"realityshows"`
	Date         time.Time `json:"date"`
	Nickname     string    `json:"nickname"`
}

type Comment struct {
	CommentID       int
	PostID          int
	Comment         string    `json:"comment"`
	CommentNickname string    `json:"commentnickname"`
	CommentDate     time.Time `json:"commentdate"`
}

func Posting(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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

	post.Date = time.Now()
	nickname, err := nicknameFromSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if post.Title != "" && post.Content != "" {
		_, err := database.Db.Exec(`INSERT INTO posts (title, content, category1, category2, category3, date, nickname) VALUES (?, ?, ?, ?, ?, ?, ?)`,
			post.Title, post.Content, post.Movies, post.Serials, post.RealityShows, post.Date, nickname)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func Commenting(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		addComment(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Error 405, method not allowed")
		return
	}
}

func addComment(w http.ResponseWriter, r *http.Request) {
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment.CommentDate = time.Now()
	id := r.URL.Query().Get("id")
	nickname, err := nicknameFromSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if comment.Comment != "" {
		_, err := database.Db.Exec(`INSERT INTO comments (comment, date, nickname, postID) VALUES (?, ?, ?, ?)`,
			comment.Comment, comment.CommentDate, nickname, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

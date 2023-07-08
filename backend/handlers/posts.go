package handlers

import (
	"encoding/json"
	"net/http"
	"real-time-forum/backend/database"
)

func DisplayPosts(w http.ResponseWriter, r *http.Request) {
	var allPosts []Post
	var post Post

	rows, err := database.Db.Query(`SELECT postID, title, content, nickname, date, category1, category2, category3 FROM posts`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Nickname, &post.Date, &post.Movies, &post.Serials, &post.RealityShows)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		allPosts = append(allPosts, post)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(allPosts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

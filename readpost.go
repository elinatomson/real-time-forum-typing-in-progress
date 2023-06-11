package main

import (
	"encoding/json"
	"net/http"
)

func readPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var post Post
	err := db.QueryRow(`SELECT postID, title, content, category1, category2, category3, nickname, date FROM posts WHERE postID=?`, id).Scan(
		&post.ID, &post.Title, &post.Content, &post.Movies, &post.Serials, &post.Realityshows, &post.Nickname, &post.Date,
	)
	if err != nil {
		panic(err)
	}

	//set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	//send the post as a JSON response
	json.NewEncoder(w).Encode(post)
}

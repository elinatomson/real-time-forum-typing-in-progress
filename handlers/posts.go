package handlers

import (
	"encoding/json"
	"net/http"
)

func DisplayPosts(w http.ResponseWriter, r *http.Request) {
	var allPosts []Post
	var post Post

	rows, err := db.Query(`SELECT postID, title, content, nickname, date, category1, category2, category3 FROM posts`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Nickname, &post.Date, &post.Movies, &post.Serials, &post.Realityshows)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		allPosts = append(allPosts, post)
	}
	//send the posts as a JSON response
	json.NewEncoder(w).Encode(allPosts)
}

//TO-DO: check if all local code is ok

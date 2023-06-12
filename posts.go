package main

import (
	"encoding/json"
	"net/http"
)

func displayPosts(w http.ResponseWriter, r *http.Request) {
	var allPosts []Post
	var post Post

	rows, _ := db.Query(`SELECT postID, title, content, nickname, date, category1, category2, category3 FROM posts`)
	for rows.Next() {
		rows.Scan(&post.ID, &post.Title, &post.Content, &post.Nickname, &post.Date, &post.Movies, &post.Serials, &post.Realityshows)
		allPosts = append(allPosts, post)
	}
	//send the posts as a JSON response
	json.NewEncoder(w).Encode(allPosts)
}

//TO-DO: check if all local code is ok

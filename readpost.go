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

	var comment Comment
	var allComments []Comment
	rows1, _ := db.Query(`SELECT commentID, postID, comment, nickname, date FROM comments WHERE postID = ?`, id)
	for rows1.Next() {
		rows1.Scan(&comment.CommentID, &comment.PostID, &comment.Comment, &comment.CommentNickname, &comment.CommentDate)
		allComments = append(allComments, comment)
	}

	//struct to hold both the post and comments
	type PostWithComments struct {
		Post     Post      `json:"post"`
		Comments []Comment `json:"comments"`
	}

	//create an instance of the struct and populate it with the post and comments
	postWithComments := PostWithComments{
		Post:     post,
		Comments: allComments,
	}

	//set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	//send the post and comments as a JSON response
	json.NewEncoder(w).Encode(postWithComments)
}

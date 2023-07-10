package handlers

import (
	"encoding/json"
	"net/http"
	"real-time-forum/backend/database"
)

func ReadPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var post Post
	rows, err := database.Db.Query(`SELECT postID, title, content, category1, category2, category3, nickname, date FROM posts WHERE postID=?`, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Movies, &post.Serials, &post.RealityShows, &post.Nickname, &post.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	var comment Comment
	var allComments []Comment
	rows1, err := database.Db.Query(`SELECT commentID, postID, comment, nickname, date FROM comments WHERE postID = ?`, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows1.Close()

	for rows1.Next() {
		err = rows1.Scan(&comment.CommentID, &comment.PostID, &comment.Comment, &comment.CommentNickname, &comment.CommentDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		allComments = append(allComments, comment)
	}

	type PostWithComments struct {
		Post     Post      `json:"post"`
		Comments []Comment `json:"comments"`
	}

	postWithComments := PostWithComments{
		Post:     post,
		Comments: allComments,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(postWithComments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func frontPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error 404, page not found")
		return
	}
	user, _ := nicknameFromSession(r)

	var allPosts []Post
	var post Post
	rows, _ := db.Query(`SELECT postID, title, nickname, date, category1, category2, category3 FROM posts ORDER BY postID DESC`)
	for rows.Next() {
		rows.Scan(&post.ID, &post.Title, &post.Nickname, &post.Date, &post.Movies, &post.Serials, &post.Realityshows)
		allPosts = append(allPosts, post)
	}
	//"Logged" for specifying in frontpage.html, what the user can do as being logged in or not
	Logged := loggedIn(w, r)
	page := Page{Posts: allPosts, LogIn: Logged, User: user}
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, page)
}

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var tpl *template.Template

func main() {
	var err error
	db, err = sql.Open("sqlite3", "data.db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//creating all tables into the database
	allDataBases()

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/userpage", userPage)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", logIn)
	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/message", messageing)
	http.HandleFunc("/messages", getMessages)
	http.HandleFunc("/logout", logOut)
	http.HandleFunc("/create-post", posting)
	http.HandleFunc("/commenting", commenting)
	http.HandleFunc("/posts", displayPosts)
	http.HandleFunc("/readpost", readPost)
	http.HandleFunc("/session", session)

	//new endpoint for live chat endpoint - maybe it should connect to handler
	http.HandleFunc("/ws", WebsocketHandler)

	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	fmt.Printf("Starting server at port 8080\nOpen http://localhost:8080\nUse Ctrl+C to close the port\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

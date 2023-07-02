package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"real-time-forum/handlers"

	_ "github.com/mattn/go-sqlite3"
)

var tpl *template.Template

func main() {
	//creating all tables into the database
	handlers.AllDataBases()

	http.HandleFunc("/", handlers.MainPage)
	http.HandleFunc("/userpage", handlers.UserPage)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.LogIn)
	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/message", handlers.Messageing)
	http.HandleFunc("/messages", handlers.GetMessages)
	http.HandleFunc("/messages/unread", handlers.UnreadMessages)
	http.HandleFunc("/messages/markAsRead", handlers.MessagesAsRead)
	http.HandleFunc("/logout", handlers.LogOut)
	http.HandleFunc("/create-post", handlers.Posting)
	http.HandleFunc("/commenting", handlers.Commenting)
	http.HandleFunc("/posts", handlers.DisplayPosts)
	http.HandleFunc("/readpost", handlers.ReadPost)
	http.HandleFunc("/session", handlers.Sessions)
	http.HandleFunc("/ws", handlers.Websocket)

	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	fmt.Printf("Starting server at port 8080\nOpen http://localhost:8080\nUse Ctrl+C to close the server\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

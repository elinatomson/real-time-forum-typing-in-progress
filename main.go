package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"real-time-forum/backend/database"
	"real-time-forum/backend/handlers"

	_ "github.com/mattn/go-sqlite3"
)

var tpl *template.Template

func main() {
	//creating all tables into the database
	database.AllDataBases()

	http.HandleFunc("/", handlers.MainPage)
	http.HandleFunc("/userpage", handlers.UserPage)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.LogIn)
	http.HandleFunc("/users", handlers.GetUsers)
	http.HandleFunc("/message", handlers.Messageing)
	http.HandleFunc("/messages", handlers.GetMessages)
	http.HandleFunc("/messages/unread", handlers.UnreadMessages)
	http.HandleFunc("/messages/mark-as-read", handlers.MessagesAsRead)
	http.HandleFunc("/logout", handlers.LogOut)
	http.HandleFunc("/create-post", handlers.Posting)
	http.HandleFunc("/commenting", handlers.Commenting)
	http.HandleFunc("/posts", handlers.DisplayPosts)
	http.HandleFunc("/read-post", handlers.ReadPost)
	http.HandleFunc("/session", handlers.Sessions)
	http.HandleFunc("/ws", handlers.Websocket)

	fileServer := http.FileServer(http.Dir("./frontend/"))
	http.Handle("/frontend/", http.StripPrefix("/frontend", fileServer))
	fmt.Printf("Starting server at port 8080\nOpen http://localhost:8080\nUse Ctrl+C to close the server\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

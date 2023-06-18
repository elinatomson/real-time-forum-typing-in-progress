package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"real-time-forum/websocket"
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

	http.HandleFunc("/", handler)
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
	http.HandleFunc("/ws", websocket.WebsocketHandler)

	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	fmt.Printf("Starting server at port 8080\nOpen http://localhost:8080\nUse Ctrl+C to close the port\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error 404, page not found")
		return
	}

	/*
		tmpl := template.Must(template.ParseFiles("static/index.html"))
		tmpl.Execute(w, nil)
	*/

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Println("Failed to parse template:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server Error")
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Failed to execute template:", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server Error")
		return
	}
}

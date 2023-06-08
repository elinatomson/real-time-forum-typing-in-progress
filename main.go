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
	AllDataBases()

	http.HandleFunc("/", handler)
	http.HandleFunc("/register", register)

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
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, nil)
}

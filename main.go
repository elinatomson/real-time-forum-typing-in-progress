package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)         

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
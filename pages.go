package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error 404, page not found")
		return
	}

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

func userPage(w http.ResponseWriter, r *http.Request) {
	//checking if the user is logged in
	logged, err := checkSession(w, r)
	if err != nil {
		fmt.Printf("Error with userPage reading checkSession: %v", err)
	}
	if logged == 0 {
		http.Redirect(w, r, "/", http.StatusFound) //redirecting to the original main page if the user is not logged in
		return
	}
}

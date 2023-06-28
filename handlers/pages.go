package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Println("Page not found")
		http.Error(w, "This page does not exist.", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		log.Println("Failed to parse template:", err)
		http.Error(w, "Something went wrong, please try again.", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("Failed to execute template:", err)
		http.Error(w, "Something went wrong, please try again.", http.StatusInternalServerError)
		return
	}
}

func UserPage(w http.ResponseWriter, r *http.Request) {
	//checking if the user is logged in
	logged, err := checkSession(w, r)
	if err != nil {
		log.Printf("Error with UserPage reading checkSession: %v", err)
		http.Error(w, "Something went wrong, please try again.", http.StatusInternalServerError)
		return
	}
	if logged == 0 {
		//redirecting to the original main page if the user is not logged in
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
}

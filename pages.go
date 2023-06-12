package main

import (
	"net/http"
)

func userPage(w http.ResponseWriter, r *http.Request) {
	//checking if the user is logged in
	logged, _ := checkSession(w, r)
	if logged == 0 {
		http.Redirect(w, r, "/", http.StatusFound) //redirecting to the original main page if the user is not logged in
		return
	}
}

//TO-DO: probably refresh sets logged == 0, check where in code refresh is handled or if there is need for pointers

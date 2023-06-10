package main

import (
	"net/http"
)

func logOut(w http.ResponseWriter, r *http.Request) {
	err := deleteSession(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

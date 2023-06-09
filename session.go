package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func addCookie(w http.ResponseWriter, nickname string) {
	//Generate a new UUID for a session.
	uuid, _ := uuid.NewV4()
	value := uuid.String()
	expire := time.Now().Add(1 * time.Hour)
	cookie := http.Cookie{
		Name:    "sessionId",
		Value:   value,
		Expires: expire,
	}
	http.SetCookie(w, &cookie)
	_, err := db.Exec(`INSERT INTO sessions (nickname, cookie)  VALUES(?, ?)	`, nickname, uuid.String())
	if err != nil {
		return
	}
}
func checkSession(w http.ResponseWriter, r *http.Request) (int, error) {
	cookie, err := r.Cookie("sessionId")
	//Checks if there's a cookie
	if err != nil {
		fmt.Println("cookie was not found")
		if err == http.ErrNoCookie {
			return 0, nil
		}
		return 0, err
	}
	//Verify that the uuID is valid
	uuid, err := uuid.FromString(cookie.Value)
	fmt.Println(cookie.Value)
	if err != nil {
		return 0, err
	}
	var nickname string
	err = db.QueryRow(`SELECT nickname FROM sessions WHERE cookie = ?`, uuid.String()).Scan(&nickname)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

// Convert checkSession into bool value, in order to use it in html
func loggedIn(w http.ResponseWriter, r *http.Request) bool {
	id, _ := checkSession(w, r)
	if id == 0 {
		return false
	}
	return true
}

// If logging out, then deleting session
func deleteSession(r *http.Request) error {
	cookie, err := r.Cookie("sessionId")
	if err != nil {
		fmt.Println("theres no cookie to be found")
		if err == http.ErrNoCookie {
			return nil
		}
		return err
	}
	uuid, err := uuid.FromString(cookie.Value)
	if err != nil {
		return err
	}
	_, err = db.Exec(`DELETE FROM sessions WHERE cookie = ?`, uuid.String())
	if err != nil {
		return err
	}
	return nil
}

// Getting nickname from session
func nicknameFromSession(r *http.Request) (string, error) {
	cookie, err := r.Cookie("sessionId")
	if err != nil {
		return "", err
	}
	uuid, err := uuid.FromString(cookie.Value)
	if err != nil {
		return "", err
	}
	var nickname string
	err = db.QueryRow(`SELECT nickname FROM sessions WHERE cookie = ?`, uuid.String()).Scan(&nickname)
	if err != nil {
		return "", err
	}
	return nickname, nil
}

func sessionExists(db *sql.DB, nickname string) bool {
	sqlStmt := `SELECT nickname FROM sessions WHERE nickname = ?`
	err := db.QueryRow(sqlStmt, nickname).Scan(&nickname)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
}

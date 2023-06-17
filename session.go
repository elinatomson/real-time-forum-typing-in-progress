package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

type Session struct {
	Cookie   string `json:"cookie"`
	Nickname string `json:"nickname"`
}

func addCookie(w http.ResponseWriter, nickname string) {
	//generate a new UUID for a session.
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

func session(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decodeSession(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Error 405, method not allowed")
		return
	}
}

func decodeSession(w http.ResponseWriter, r *http.Request) {
	var session Session
	err := json.NewDecoder(r.Body).Decode(&session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = getCookieFromSession(w, session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getCookieFromSession(w http.ResponseWriter, session Session) error {
	// Query the database to retrieve the stored cookie value
	stmt := `SELECT cookie FROM sessions WHERE cookie = ?`
	row := db.QueryRow(stmt, session.Cookie)
	var storedCookie string
	err := row.Scan(&storedCookie)
	if err != nil {
		fmt.Fprintf(w, "Cookie does not match!")
	}

	fmt.Println(session.Cookie)
	fmt.Println(storedCookie)

	// Compare the client-side cookie with the stored cookie
	if session.Cookie == storedCookie {
		fmt.Fprintf(w, "Cookie matches!")
	} else {
		fmt.Fprintf(w, "Cookie does not match!")
	}
	return nil
}

func checkSession(w http.ResponseWriter, r *http.Request) (int, error) {
	cookie, err := r.Cookie("sessionId")
	//checks if there's a cookie
	if err != nil {
		fmt.Println("cookie was not found")
		if err == http.ErrNoCookie {
			return 0, nil
		}
		return 0, err
	}
	//verify that the uuID is valid
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

// convert checkSession into bool value, in order to use it in html
func loggedIn(w http.ResponseWriter, r *http.Request) bool {
	id, _ := checkSession(w, r)
	if id == 0 {
		return false
	}
	return true
}

// if logging out, then deleting session
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

// getting nickname from session
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

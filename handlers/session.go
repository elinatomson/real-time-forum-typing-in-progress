package handlers

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
		log.Println(err)
	}
}

func Sessions(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
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
		log.Println(err)
		fmt.Fprintf(w, "Cookie does not match!")
		return nil
	}
	// Compare the client-side cookie with the stored cookie
	if session.Cookie == storedCookie {
		fmt.Fprintf(w, "Cookie matches!")
	} else {
		fmt.Fprintf(w, "Cookie does not match!")
	}
	return nil
}

// if logging out, then deleting session
func deleteSession(r *http.Request) error {
	cookie, err := r.Cookie("sessionId")
	if err != nil {
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
	stmt := `SELECT nickname FROM sessions WHERE cookie = ?`
	row := db.QueryRow(stmt, uuid.String())
	var nickname string
	err = row.Scan(&nickname)
	if err != nil {
		return "", err
	}

	return nickname, nil
}

func sessionExists(db *sql.DB, nickname string) bool {
	stmt := `SELECT nickname FROM sessions WHERE nickname = ?`
	row := db.QueryRow(stmt, nickname)
	var storedNickname string
	err := row.Scan(&storedNickname)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
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
	stmt := `SELECT nickname FROM sessions WHERE cookie = ?`
	row := db.QueryRow(stmt, uuid.String())
	var nickname string

	err = row.Scan(&nickname)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

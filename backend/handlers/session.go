package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"real-time-forum/backend/database"
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
	_, err := database.Db.Exec(`INSERT INTO sessions (nickname, cookie)  VALUES(?, ?)	`, nickname, uuid.String())
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
	stmt := `SELECT cookie FROM sessions WHERE cookie = ?`
	row := database.Db.QueryRow(stmt, session.Cookie)
	var storedCookie string
	err := row.Scan(&storedCookie)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Cookie does not match!")
		return nil
	}

	if session.Cookie == storedCookie {
		fmt.Fprintf(w, "Cookie matches!")
	} else {
		fmt.Fprintf(w, "Cookie does not match!")
	}
	return nil
}

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
	_, err = database.Db.Exec(`DELETE FROM sessions WHERE cookie = ?`, uuid.String())
	if err != nil {
		return err
	}
	return nil
}

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
	row := database.Db.QueryRow(stmt, uuid.String())
	var nickname string
	err = row.Scan(&nickname)
	if err != nil {
		return "", err
	}

	return nickname, nil
}

func sessionExists(db *sql.DB, nickname string) bool {
	stmt := `SELECT nickname FROM sessions WHERE nickname = ?`
	row := database.Db.QueryRow(stmt, nickname)
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

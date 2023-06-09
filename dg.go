package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"
)

var db *sql.DB

const UserDB = `
	CREATE TABLE IF NOT EXISTS users (
		ID					INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		nickname			TEXT,
		age					INTEGER,
		gender				TEXT,
		firstname			TEXT,
		lastname			TEXT,
		email 				TEXT,
		password			TEXT
	)`
const PostDB = `
	CREATE TABLE IF NOT EXISTS posts (
		postID				INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		title   			TEXT,
		content   			TEXT,
		category1			TEXT,
		category2			TEXT,
		category3			TEXT,
		date				DATETIME,
		nickname			TEXT NOT NULL
	)`
const CommentDB = `
	CREATE TABLE IF NOT EXISTS comments (
		commentID			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		comment   			TEXT,
		date       			DATETIME,
		postID            	INTEGER,
		postTitle   		TEXT,
		nickname          	TEXT
		)`
const SessionDB = `
	CREATE TABLE IF NOT EXISTS sessions (
		nickname 			TEXT NOT NULL,
		cookie				TEXT NOT NULL
	)`

type Page struct {
	Posts    []Post
	Comments []Comment
	ID       int
	LogIn    bool
	User     string
}

type Comment struct {
	CommentID       int
	PostID          int
	PostTitle       string
	Comment         string
	CommentNickname string
	CommentDate     time.Time
}

type Session struct {
	Cookie   string
	Nickname string
}

func CreateDataBase(data string) error {
	stmt, err := db.Prepare(data)
	if err != nil {
		return err
	}
	stmt.Exec()
	stmt.Close()
	return nil
}
func AllDataBases() {
	//create named tables
	CreateDataBase(UserDB)
	CreateDataBase(PostDB)
	CreateDataBase(CommentDB)
	CreateDataBase(SessionDB)
}
func TitleFromPost(r *http.Request) (string, error) {
	id := r.FormValue("id")
	id1, _ := strconv.Atoi(id)
	var title string
	err := db.QueryRow(`SELECT title FROM posts WHERE postID = ?`, id1).Scan(&title)
	if err != nil {
		return "", err
	}
	return title, nil
}

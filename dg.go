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
		username			TEXT
	)`
const CommentDB = `
	CREATE TABLE IF NOT EXISTS comments (
		commentID			INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		comment   			TEXT,
		likes      			INTEGER,
		dislikes   			INTEGER,
		date       			DATETIME,
		postID            	INTEGER,
		postTitle   		TEXT,
		username          	INTEGER
		)`
const PostLikeDB = `
	CREATE TABLE IF NOT EXISTS postlikes (
		username 		TEXT NOT NULL,
		postID 			INTEGER,
		postTitle		TEXT
		)`
const PostDisLikeDB = `
		CREATE TABLE IF NOT EXISTS postdislikes (
		username		TEXT NOT NULL,
		postID 			INTEGER,
		postTitle		TEXT
		)`
const CommentLikeDB = `
	CREATE TABLE IF NOT EXISTS commentlikes (
		username 		TEXT NOT NULL,
		commentID 		INTEGER,
		like 			INTEGER
		)`
const CommentdisLikeDB = `
	CREATE TABLE IF NOT EXISTS commentdislikes (
		username 		TEXT NOT NULL,
		commentID 		INTEGER,
		dislike 		INTEGER
		)`
const SessionDB = `
	CREATE TABLE IF NOT EXISTS sessions (
		username 			TEXT NOT NULL,
		cookie				TEXT NOT NULL
	)`

type Page struct {
	Posts    []Post
	Comments []Comment
	Likes    []Like
	DisLikes []DisLike
	ID       int
	Like     int
	DisLike  int
	LogIn    bool
	User     string
}
type Post struct {
	ID         int
	Title      string
	Content    string
	Category1  string
	Category2  string
	Category3  string
	Username   string
	Date       time.Time
	Comments   []Comment
	Likes      []Like
	LikesNb    int
	DisLikesNb int
}
type Comment struct {
	CommentID       int
	PostID          int
	PostTitle       string
	Comment         string
	CommentUsername string
	CommentDate     time.Time
	CommentLikes    int
	CommentDislikes int
}
type Like struct {
	Username  string
	PostID    int
	PostTitle string
}
type DisLike struct {
	Username  string
	PostID    int
	PostTitle string
}
type CommentLike struct {
	Username  string
	PostID    int
	PostTitle string
}
type CommentDislike struct {
	Username  string
	PostID    int
	PostTitle string
}
type Session struct {
	Cookie   string
	Username string
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
	CreateDataBase(PostLikeDB)
	CreateDataBase(PostDisLikeDB)
	CreateDataBase(CommentLikeDB)
	CreateDataBase(CommentdisLikeDB)
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

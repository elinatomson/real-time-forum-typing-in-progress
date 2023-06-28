package handlers

import (
	"database/sql"
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
		password			TEXT,
		last_message_date   TIMESTAMP DEFAULT NULL
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
		nickname          	TEXT
		)`
const SessionDB = `
	CREATE TABLE IF NOT EXISTS sessions (
		nickname 			TEXT NOT NULL,
		cookie				TEXT NOT NULL
	)`

// integer default 0 means that the message is unread. after reading it it becomes 1 indicating as read
const MessageDB = `
	CREATE TABLE IF NOT EXISTS messages (
		messageID				INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		message 			TEXT,
		nicknamefrom		TEXT,
		nicknameto			TEXT,
		date       			DATETIME,
		read          		INTEGER DEFAULT 0 
)`

func createDataBaseTable(data string) error {
	stmt, err := db.Prepare(data)
	if err != nil {
		return err
	}
	stmt.Exec()
	stmt.Close()
	return nil
}
func AllDataBases() {
	var err error
	db, err = sql.Open("sqlite3", "data.db")
	if err != nil {
		panic(err.Error())
	}
	//create named tables
	createDataBaseTable(UserDB)
	createDataBaseTable(PostDB)
	createDataBaseTable(CommentDB)
	createDataBaseTable(SessionDB)
	createDataBaseTable(MessageDB)
}

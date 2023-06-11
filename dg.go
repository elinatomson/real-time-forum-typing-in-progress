package main

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
		nickname          	TEXT
		)`
const SessionDB = `
	CREATE TABLE IF NOT EXISTS sessions (
		nickname 			TEXT NOT NULL,
		cookie				TEXT NOT NULL
	)`

func createDataBase(data string) error {
	stmt, err := db.Prepare(data)
	if err != nil {
		return err
	}
	stmt.Exec()
	stmt.Close()
	return nil
}
func allDataBases() {
	//create named tables
	createDataBase(UserDB)
	createDataBase(PostDB)
	createDataBase(CommentDB)
	createDataBase(SessionDB)
}

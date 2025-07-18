package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	DB = db

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY NOT NULL,
		email TEXT NOT NULL UNIQUE,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (	
		id INTEGER PRIMARY KEY AUTOINCREMENT,
	    event_id INTEGER NOT NULL,
	    user_id TEXT NOT NULL,
	    FOREIGN KEY(event_id) REFERENCES events(id),
	    FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic(err)
	}
}

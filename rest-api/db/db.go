package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func InitDB() {
	var err error

	Db, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("failed to connect database")
	}

	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
		    id INTEGER PRIMARY KEY AUTOINCREMENT,
		    email TEXT UNIQUE NOT NULL,
		    password TEXT NOT NULL
		)
	`

	_, err := Db.Exec(createUsersTable)

	if err != nil {
		panic("failed to create users table " + err.Error())
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
		    Id INTEGER PRIMARY KEY AUTOINCREMENT,
			Name TEXT UNIQUE NOT NULL,
		    Description TEXT NOT NULL,
		    Location TEXT NOT NULL,
		    DateTime DATETIME NOT NULL,
		    UserId INTEGER NOT NULL
		)
	`

	_, err = Db.Exec(createEventsTable)

	if err != nil {
		panic("failed to create events table " + err.Error())
	}
}

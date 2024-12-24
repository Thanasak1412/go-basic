package models

import (
	"database/sql"
	"fmt"
	"github.com/Thanasak1412/go-rest-api/db"
	"time"
)

type Event struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserId      int64     `json:"userId"`
}

func (e *Event) Save() error {
	stmt, err := db.Db.Prepare(`
		INSERT INTO events(name, description, location, dateTime, userId)
		VALUES (?, ?, ?, ?, ?)
	`)

	if err != nil {
		return err
	}

	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)

	event, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}
	lastInsertId, err := event.LastInsertId()

	if err != nil {
		return err
	}

	e.Id = lastInsertId

	return nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT id, name, description, location, dateTime FROM events WHERE id = ?"
	row := db.Db.QueryRow(query, id)

	var event Event
	if err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime); err != nil {
		fmt.Println("error scanning event", err.Error())

		return nil, err
	}

	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT id, name, description, location, datetime FROM events"

	rows, err := db.Db.Query(query)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var events []Event

	for rows.Next() {
		var event Event

		if err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (e *Event) UpdateEvent() error {
	query := "UPDATE events SET name = ?, description = ?, location = ? WHERE id = ?"

	stmt, err := db.Db.Prepare(query)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Id)

	if err != nil {
		return err
	}

	_, err = result.LastInsertId()

	if err != nil {
		return err
	}

	return nil
}

func DeleteEventById(id int64) error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.Db.Prepare(query)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(id)

	if err != nil {
		return err
	}

	_, err = result.RowsAffected()

	if err != nil {
		return err
	}

	return nil
}

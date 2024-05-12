package models

import (
	"time"

	"example.com/REST-API-Event-Booking/db"
)

type Event struct {
	ID          int64
	Name        string     `binding:"required"`
	Description string     `binding:"required"`
	Location    string     `binding:"required"`
	DateTime    *time.Time `binding:"required"`
	UserID      int64
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return nil
}

func (event *Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func (event *Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	_, err := db.DB.Exec(query, event.ID)
	return err
}

func (e *Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	_, err := db.DB.Exec(query, e.ID, userId)
	return err
}

func (e *Event) CancelRegistration(eventId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ?"
	_, err := db.DB.Exec(query, eventId)
	return err
}

// type Event struct {
// 	ID          int64      `db:"id" binding:"required"`
// 	Name        string     `db:"name" binding:"required"`
// 	Description string     `db:"description" binding:"required"`
// 	Location    string     `db:"location" binding:"required"`
// 	DateTime    *time.Time `db:"dateTime" binding:"required"`
// 	UserID      int        `db:"user_id"`
// }
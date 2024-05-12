package models

import (
	"database/sql"
	"errors"
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

type Registration struct {
	ID      int64
	EventID int64
	UserID  int64
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
	exists, err := e.CheckRegistrationExists(userId)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("Event registration already exists")
	}

	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"
	_, err = db.DB.Exec(query, e.ID, userId)
	return err
}

func (e *Event) CheckRegistrationExists(userId int64) (bool, error) {
	query := "SELECT COUNT(*) FROM registrations WHERE event_id = ? AND user_id = ?"
	var count int
	err := db.DB.QueryRow(query, e.ID, userId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func GetEventRegistrationById(id int64) (*Registration, error) {
	query := "SELECT * FROM registrations WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var registration Registration
	err := row.Scan(&registration.ID, &registration.EventID, &registration.UserID)
	if err != nil {
		return nil, err
	}

	return &registration, nil
}

func CancelEventRegistration(eventId int64) error {
	_, err := GetEventRegistrationById(eventId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Event registration not found")
		}
		return err
	}

	query := "DELETE FROM registrations WHERE id = ?"
	_, err = db.DB.Exec(query, eventId)
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

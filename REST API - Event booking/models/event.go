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

// type Event struct {
// 	ID          int64      `db:"id" binding:"required"`
// 	Name        string     `db:"name" binding:"required"`
// 	Description string     `db:"description" binding:"required"`
// 	Location    string     `db:"location" binding:"required"`
// 	DateTime    *time.Time `db:"dateTime" binding:"required"`
// 	UserID      int        `db:"user_id"`
// }

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
	return err
}

package models

import (
	"database/sql"
	"errors"

	"example.com/REST-API-Event-Booking/db"
	"example.com/REST-API-Event-Booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func FindUserByEmail(email string) (*User, error) {
	var user User
	query := "SELECT id, email, password FROM users WHERE email = ?"
	err := db.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}


func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	row, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer row.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := row.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Invalid credentials" + err.Error())
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Invalid credentials" + err.Error())
	}

	return nil
}

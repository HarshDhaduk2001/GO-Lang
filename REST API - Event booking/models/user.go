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
		return nil, errors.New("failed to find user by email")
	}
	return &user, nil
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	row, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("failed to prepare database for user insertion")
	}
	defer row.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return errors.New("failed to hash user password")
	}

	result, err := row.Exec(u.Email, hashedPassword)
	if err != nil {
		return errors.New("failed to insert user into database")
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return errors.New("failed to retrieve last inserted ID")
	}

	u.ID = userId
	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}

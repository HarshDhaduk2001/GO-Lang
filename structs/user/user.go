package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First name, Last name and birthday are required fields.")
	}

	return &User{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}, nil
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			"ADMIN",
			"ADMIN",
			"---",
			time.Now(),
		},
	}
}

func (u User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createdAt)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

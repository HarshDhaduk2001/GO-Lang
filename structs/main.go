package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter first name: ")
	lastName := getUserData("Please enter last name: ")
	birthdate := getUserData("Please enter birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.NewUser(firstName, lastName, birthdate)
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("admin@yopmail.com", "password")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(propmptText string) string {
	fmt.Print(propmptText)
	var value string
	fmt.Scanln(&value)
	return value
}

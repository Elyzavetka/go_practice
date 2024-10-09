package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func (user User) outputUserDetails() {
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

func (user *User) clearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func newUser(firstName, lastName, birthdate string) *User {
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")


	appUser := newUser(userFirstName, userLastName, userBirthdate)

	

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

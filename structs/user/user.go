package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func (user User) OutputUserDetails() {
	fmt.Println(user.firstName, user.lastName, user.birthDate)
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func NewUser(firstName, lastName, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName, lastName, birthdate are required.")
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil
}

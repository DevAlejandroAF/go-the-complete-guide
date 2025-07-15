package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	User
	email    string
	password string
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("[!] Fields can't be empty.")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(firstName, lastName, birthDate, email, password string) (*Admin, error) {
	if firstName == "" || lastName == "" || birthDate == "" || email == "" || password == "" {
		return nil, errors.New("[!] Fields can't be empty.")
	}

	return &Admin{
		User: User{
			firstName: firstName,
			lastName:  lastName,
			birthDate: birthDate,
			createdAt: time.Now(),
		},
		email:    email,
		password: password,
	}, nil
}

func (u User) OutputUserDetails() {
	fmt.Println("User Form:")
	fmt.Println(u.firstName, u.lastName, u.birthDate)
	fmt.Println()
}

func (a Admin) OutputAdminDetails() {
	fmt.Println("Admin Form:")
	fmt.Println(a.firstName, a.lastName, a.birthDate, a.email)
	fmt.Println()
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

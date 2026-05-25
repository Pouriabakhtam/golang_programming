package main

import (
	"fmt"
)

type AppUser struct {
	FirstName   string
	LastName    string
	DateOfBirth string
}

func main() {
	UserFirstName := GetUserInput("Please enter your first name: ")
	UserLastName := GetUserInput("Please enter your last name: ")
	UserDateOfBirth := GetUserInput("Please enter your date of birth: ")

	User := AppUser{
		FirstName:   UserFirstName,
		LastName:    UserLastName,
		DateOfBirth: UserDateOfBirth,
	}

	GetOutput(User)
}

func GetUserInput(p string) string {
	var input string
	fmt.Print(p)
	_, err := fmt.Scan(&input)
	if err != nil {
		panic("Error reading user input. Please enter a valid string.")
	}
	return input
}

func GetOutput(u AppUser) {
	fmt.Printf("Welcome to the app, please confirm your firstname is \033[1;36m%s\033[0m, last name is \033[1;32m%s\033[0m, and date of birth is \033[1;35m%s\033[0m.\n",
		u.FirstName,
		u.LastName,
		u.DateOfBirth)
}

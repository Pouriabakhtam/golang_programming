package main

import (
	"fmt"
	"strings"

	"struct.com/user"
)

func main() {
	Userlevelaccess := GetUserAcessLevel()
	switch Userlevelaccess {
	case "user":
		UserFirstName := GetUserInput("Please enter your first name: ")
		UserLastName := GetUserInput("Please enter your last name: ")
		UserDateOfBirth := GetUserInput("Please enter your date of birth: ")
		User, err := user.NewUser(UserFirstName, UserLastName, UserDateOfBirth)
		if err != nil {
			panic(err)
		}
		User.GetOutput()
		User.SwapFirstAndLastName()
		User.GetOutput()
	case "admin":
		UserFirstName := GetUserInput("Please enter your first name: ")
		UserLastName := GetUserInput("Please enter your last name: ")
		UserDateOfBirth := GetUserInput("Please enter your date of birth: ")
		AdminEmail := GetUserInput("Please enter your email: ")
		AdminPhone := GetUserInput("Please enter your phone number: ")

		Admin, err := user.NewAdminUser(UserFirstName, UserLastName, UserDateOfBirth, AdminEmail, AdminPhone)
		if err != nil {
			panic(err)
		}
		Admin.GetOutput()
		Admin.GetAdminOutput()
		Admin.SwapFirstAndLastName()
		Admin.GetOutput()
		Admin.GetAdminOutput()
	}

}

func GetUserInput(p string) string {
	var input string
	fmt.Print(p)
	// _, err := fmt.Scan(&input)
	// _, err := fmt.Scanln(&input)
	fmt.Scanln(&input)
	return input
}

func GetUserAcessLevel() string {
	var input string
	fmt.Print("Please enter your access level (admin/user): ")
	fmt.Scan(&input)
	return strings.ToLower(input)
}

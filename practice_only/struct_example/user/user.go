package user

import (
	"errors"
	"fmt"
	"time"
)

type AppUser struct {
	firstName   string
	lastName    string
	dateOfBirth string
	createdAt   time.Time
}

type AdminUser struct {
	email       string
	phoneNumber string
	AppUser     // embedding AppUser struct into AdminUser struct
}

// struct methods
func (u *AppUser) GetOutput() {
	fmt.Printf("Welcome to the app, please confirm your firstname is \033[1;36m%s\033[0m, last name is \033[1;32m%s\033[0m, and date of birth is \033[1;35m%s\033[0m, and created at is \033[1;34m%s\033[0m.\n",
		(*u).firstName,
		(*u).lastName,
		(*u).dateOfBirth,
		(*u).createdAt.Format("2006-01-02 15:04:05"),
	)
}

func (a *AdminUser) GetAdminOutput() {
	fmt.Printf("You have admin access, here is your email: \033[1;31m%s\033[0m and phone number: \033[1;33m%s\033[0m.\n",
		a.email,
		a.phoneNumber,
	)
}

// struct method
func (u *AppUser) SwapFirstAndLastName() {
	u.firstName, u.lastName = u.lastName, u.firstName
}

// struct constructor function
func NewUser(fname, lname, dob string) (*AppUser, error) {
	if fname == "" || lname == "" || dob == "" {
		return nil, errors.New("All fields must be provided and non-empty")
	}
	return &AppUser{
		firstName:   fname,
		lastName:    lname,
		dateOfBirth: dob,
		createdAt:   time.Now(),
	}, nil
}

// struct embedding
func NewAdminUser(fname, lname, dob, email, phone string) (*AdminUser, error) {
	if fname == "" || lname == "" || dob == "" || email == "" || phone == "" {
		return nil, errors.New("All fields must be provided and non-empty")
	}
	return &AdminUser{
		email:       email,
		phoneNumber: phone,
		AppUser: AppUser{
			firstName:   fname,
			lastName:    lname,
			dateOfBirth: dob,
			createdAt:   time.Now(),
		},
	}, nil
}

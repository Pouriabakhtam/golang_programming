package modules

import (
	"errors"
	"log"

	"restapi.ca/db"
	"restapi.ca/utils"
	hash "restapi.ca/utils"
)

type Users struct {
	Id       int
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u Users) Save() error {
	query := `
	INSERT INTO USER (email, password) VALUES (? , ?)
	`
	smtm, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer smtm.Close()
	result, err := smtm.Exec(u.Email, hash.HashPassword(u.Password))
	if err != nil {
		log.Fatal(err)
		return err
	}
	user_id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return err
	}
	u.Id = int(user_id)
	return err
}

func (u *Users) ValidateCredentials() error {
	query := `
	SELECT ID,PASSWORD FROM USER WHERE EMAIL=? 
	`
	row := db.DB.QueryRow(query, u.Email)
	var retrievedpassword string
	err := row.Scan(&u.Id, &retrievedpassword)
	if err != nil {
		log.Fatal("Couldn't find the user with", u.Email)
		return err
	}
	Passwordisvalid := utils.CheckPasswordHash(u.Password, retrievedpassword)
	if !Passwordisvalid {
		return errors.New("Credential is not valid")
	}

	return nil
}

package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hpass, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hpass)
}

func CheckPasswordHash(password, hashpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	return err == nil
}

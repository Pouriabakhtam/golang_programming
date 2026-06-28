package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "secretkey"

func GenerateToken(email string, user_id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   email,
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretkey))
}

func VerifyingToken(token string) (int64, error) {
	token = strings.TrimSpace(token)
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer "))
	}
	if token == "" {
		return 0, errors.New("token is empty")
	}

	parsedtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretkey), nil
	})
	if err != nil {
		return 0, fmt.Errorf("couldn't parse the token: %w", err)
	}
	if !parsedtoken.Valid {
		return 0, errors.New("token is not valid")
	}

	claim, ok := parsedtoken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userIDFloat, ok := claim["user_id"].(float64)
	if !ok {
		return 0, errors.New("user_id claim is missing or invalid")
	}

	return int64(userIDFloat), nil
}

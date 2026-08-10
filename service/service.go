package service

import (
	"auth-api/db"
	"auth-api/model"
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func Singup(database *sql.DB, user model.User) error {
	err := db.CreateUser(database, user)

	if err != nil {
		return err
	}
	fmt.Println("User saved successfully")
	return nil
}

func Login(loginRequest model.LoginRequest) (string, error) { // login service return token+err
	fmt.Println("Entered in Login request")
	user, ok := db.Users[loginRequest.Email] // browser send loginReq and db search true/false
	if !ok {
		return "", errors.New("User not found")
	}
	if user.Password != loginRequest.Password {
		return "", errors.New("invalid password")
	}

	token, err := GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Generate Token
func GenerateToken(email string) (string, error) {
	fmt.Println("Entered in Generate Token....")

	// library create token, alogrithm for sign, in token what kind of data
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})

	// token is generated not signing , now sign the token, hardcoded my-secret-key
	tokenString, err := token.SignedString([]byte("my-secret-key"))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

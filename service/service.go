package service

import (
	"auth-api/db"
	"auth-api/model"
	"errors"
	"fmt"
)

func Singup(user model.User) error {
	fmt.Println("Service called signup")
	_, ok := db.Users[user.Email]
	if ok {
		return errors.New("user already exists")
	}
	db.Users[user.Email] = user
	fmt.Println("User saved successfully")

	return nil
}

func Login(loginRequest model.LoginRequest) error {
	fmt.Println("Entered in Login request")
	user, ok := db.Users[loginRequest.Email] // browser send loginReq and db search true/false
	if !ok {
		return errors.New("User not found")
	}
	if user.Password != loginRequest.Password {
		return errors.New("invalid password")
	}
	return nil
}

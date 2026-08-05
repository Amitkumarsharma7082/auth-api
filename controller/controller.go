package controller

import (
	"auth-api/model"
	"auth-api/service"
	"encoding/json"
	"fmt"
	"net/http"
)

func SingupController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered in singup controller")

	// only post method
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	// read json body
	var user model.User

	// json->struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = service.Singup(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json") // type json
	json.NewEncoder(w).Encode("Signup Successful")     // message : successful
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Entered in login controller")

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusBadRequest)
		return
	}

	// read json
	var loginRequest model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest) // decoder ignore "name"

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	token, err := service.Login(loginRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// sending response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message" : "Login Successful",
		"token" : token,
	})
}

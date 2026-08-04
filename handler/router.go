package handler

import (
	"auth-api/controller"
	"fmt"
	"net/http"
)

func RegisterRoutes() {
	fmt.Println("Register route")
	http.HandleFunc("/signup", controller.SingupController)
	http.HandleFunc("/login", controller.LoginController)
}

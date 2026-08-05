package handler

import (
	"auth-api/controller"
	"auth-api/middleware"
	"fmt"
	"net/http"
)

func RegisterRoutes() {
	fmt.Println("Register route")
	http.HandleFunc("/signup", controller.SingupController)
	http.HandleFunc("/login", controller.LoginController)
	http.HandleFunc("/profile",
		middleware.AuthMiddleware(controller.ProfileController),
	)
}

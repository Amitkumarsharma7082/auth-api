package handler

import (
	"auth-api/controller"
	"auth-api/middleware"
	"database/sql"
	"fmt"
	"net/http"
)

func RegisterRoutes(database *sql.DB) {
	fmt.Println("Register route")
	http.HandleFunc("/signup", controller.SingupController(database))
	http.HandleFunc("/login", controller.LoginController)
	http.HandleFunc("/profile",
		middleware.AuthMiddleware(controller.ProfileController),
	)
}

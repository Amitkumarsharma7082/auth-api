package handler

import (
	"auth-api/controller"
	"auth-api/middleware"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, database *sql.DB) {
	fmt.Println("Register route")
	router.POST("/signup", controller.SingupController(database))
	router.POST("/login", controller.LoginController(database))
	router.GET("/profile", middleware.AuthMiddleware, controller.ProfileController)
	// http.HandleFunc("/profile",
	// 	middleware.AuthMiddleware(controller.ProfileController),
	// )
}

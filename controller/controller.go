package controller

import (
	"auth-api/model"
	"auth-api/service"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SingupController(database *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Entered in singup controller")

		// read json body
		var user model.User

		// json->struct
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid JSON",
			})
			return
		}

		err = service.Singup(database, user)
		c.JSON(200, gin.H{
			"message": "Signup Successful",
		})
	}
}

func LoginController(database *sql.DB) gin.HandlerFunc {
	fmt.Println("Entered in login controller")
	return func(c *gin.Context) {
		// read json
		var loginRequest model.LoginRequest
		err := c.ShouldBindJSON(&loginRequest) // decoder ignore "name"

		if err != nil {
			c.JSON(400, gin.H{
				"error": "BAD Request",
			})
			return
		}
		token, err := service.Login(database, loginRequest)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		// sending response
		c.JSON(200, gin.H{
			"message": "Login Successful",
			"token":   token,
		})
	}
}

func ProfileController(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		c.JSON(401, gin.H{
			"error": "User email not found",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Profile accessed successfully",
		"email":   email,
	})
}

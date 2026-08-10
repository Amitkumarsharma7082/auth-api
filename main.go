package main

import (
	"auth-api/db"
	"auth-api/handler"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()
	fmt.Println("Database connected")
	router := gin.Default()

	handler.RegisterRoutes(router, database)
	fmt.Println("Starting Server on :8080")

	err = router.Run(":8080")
	if err != nil {
		fmt.Println("Server Error:", err)
	}
}

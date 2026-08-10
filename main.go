package main

import (
	"auth-api/db"
	"auth-api/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {

	database, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	fmt.Println("Database connected")

	handler.RegisterRoutes(database)

	fmt.Println("Starting Server on :8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Error:", err)
	}
}

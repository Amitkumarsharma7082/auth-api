package main

import (
	"auth-api/handler"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Auth-API server")

	handler.RegisterRoutes() // register route
	fmt.Println("Starting Server on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

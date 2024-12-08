package main

import (
	"fmt"
	"zocket-api/api"
	"zocket-api/db"
)

func main() {
	fmt.Println("Bravo, Product Management Application is running")

	// Connect to MongoDB
	db.Connect()

	// Setup API routes
	api.SetupRoutes()
}

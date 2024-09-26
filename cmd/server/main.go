package main

import (
	"demoProject/config"
	"demoProject/routes"
	"demoProject/utils"
	"log"
	"net/http"
)

func main() {

	// Initialize configuration and database
	config.LoadConfig()
	utils.InitDB()

	// Initialize router
	router := routes.InitRoutes()

	// Start server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

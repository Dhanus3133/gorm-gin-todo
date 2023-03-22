package main

import (
	"github.com/Dhanus3133/gorm-gin-todo/db"
	"github.com/Dhanus3133/gorm-gin-todo/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db.Init()

	// Create a new Gin router
	router := gin.Default()

	// Set up the routes
	routes.TodoRoutes(router.Group("/api"))

	// Start the server
	router.Run(":8080")
}

package main

import (
	"project/config"
	"project/routes"

	_ "project/docs" // Swaggo docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Initialize database
	db := config.SetupDatabase()

	// Run migrations
	config.RunMigrations(db)

	// Initialize Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r, db)

	// Swagger documentation endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start the server
	r.Run(":8080")
}

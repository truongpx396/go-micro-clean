package main

import (
	"go-micro-clean/config"
	"go-micro-clean/routes"

	_ "go-micro-clean/docs" // Swaggo docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Project API
// @version         1.0
// @description     This is a sample server for a go-micro-clean.
// @termsOfService  http://example.com/terms/

// @contact.name   API Support
// @contact.url    http://www.example.com/support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @tag.name  Items
// @tag.description  Operations about items

// swag init -g cmd/main.go -o docs

func main() {
	// Initialize database
	db := config.SetupDatabase()

	db = db.Debug()

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

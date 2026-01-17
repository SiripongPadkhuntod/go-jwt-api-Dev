package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go-jwt-api/docs"

	"go-jwt-api/config"
	"go-jwt-api/routes"
)

// @title Go JWT API
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	// ===============================
	// Load ENV
	// ===============================
	config.LoadEnv()

	// ===============================
	// Connect Database
	// ===============================
	config.ConnectDB()

	// ป้องกันกรณี DB ยัง nil
	if config.DB == nil {
		
		log.Fatal("Database connection is nil")
	}

	// ===============================
	// Auto Migrate
	// ===============================
	config.AutoMigrate()

	// ===============================
	// Gin setup
	// ===============================
	r := gin.Default()

	// Routes
	routes.Setup(r)

	// Swagger (dev only)
	if os.Getenv("APP_ENV") != "prod" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Static files
	r.Static("/uploads", "./uploads")

	// ===============================
	// Run server
	// ===============================
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)
	r.Run(":" + port)

	//log swagger URL
	log.Println("Swagger UI available at http://localhost:" + port + "/swagger/index.html")
}

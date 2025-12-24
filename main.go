package main

import (
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go-jwt-api/docs"

	"go-jwt-api/config"
	"go-jwt-api/models"
	"go-jwt-api/routes"
	"go-jwt-api/seed"
)

// @title Go JWT API
// @version 1.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	// Load env & connect DB
	config.LoadEnv()
	config.ConnectDB()

	// Auto migrate
	config.DB.AutoMigrate(
		&models.User{},
		&models.Item{},
	)

	// Seed data (dev only)
	if os.Getenv("APP_ENV") == "dev" {
		seed.SeedItems()
	}

	// Routes
	routes.Setup(r)

	// Swagger
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Swagger (เปิดเฉพาะ dev)
	if os.Getenv("APP_ENV") != "prod" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// Static files (uploads)
	r.Static("/uploads", "./uploads")

	// Port
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Run server (ต้องท้ายสุด)
	r.Run(":" + port)
}

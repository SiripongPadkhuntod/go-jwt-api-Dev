package main

import (
	"os"

	"fmt"

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

	config.LoadEnv()
	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{}, &models.Item{})

	if os.Getenv("APP_ENV") == "dev" {
		seed.SeedItems()
	}

	routes.Setup(r)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + port)

	fmt.Println("Server running on port " + port)
	fmt.Println("Swagger UI: http://localhost:" + port + "/swagger/index.html")
}

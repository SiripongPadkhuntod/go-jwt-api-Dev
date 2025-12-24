package routes

import (
	"github.com/gin-gonic/gin"
	"go-jwt-api/handlers"
	"go-jwt-api/middleware"
)

func Setup(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	item := r.Group("/items")
	item.Use(middleware.AuthMiddleware())
	{
		item.GET("", handlers.GetItems)
		item.POST("", handlers.CreateItem)
		item.PUT("/:id", handlers.UpdateItem)
	}
}

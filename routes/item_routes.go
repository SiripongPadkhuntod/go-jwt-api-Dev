package routes

import (
	"github.com/gin-gonic/gin"

	"go-jwt-api/handlers"
	"go-jwt-api/middleware"
)

// SetupItemRoutes sets item related routes
func SetupItemRoutes(r *gin.Engine) {
	items := r.Group("/items")
	items.Use(middleware.AuthMiddleware())
	{
		items.GET("", handlers.GetItems)
		items.POST("", handlers.CreateItem)
		items.PUT("/:id", handlers.UpdateItem)

		// upload image
		items.POST("/:id/upload", handlers.UploadItemImage)
	}
}

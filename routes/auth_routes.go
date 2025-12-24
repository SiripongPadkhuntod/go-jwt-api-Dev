package routes

import (
	"github.com/gin-gonic/gin"

	"go-jwt-api/handlers"
)

// SetupAuthRoutes sets auth related routes
func SetupAuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}
}

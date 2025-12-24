package routes

import "github.com/gin-gonic/gin"

// Setup registers all routes
func Setup(r *gin.Engine) {
	SetupAuthRoutes(r)
	SetupItemRoutes(r)
}

package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"go-jwt-api/config"
	"go-jwt-api/models"
	"go-jwt-api/utils"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary Register
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body AuthRequest true "Register"
// @Success 200
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var req AuthRequest
	c.BindJSON(&req)

	hash, _ := utils.HashPassword(req.Password)

	user := models.User{
		Username: req.Username,
		Password: hash,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registered"})
}

// @Summary Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body AuthRequest true "Login"
// @Success 200 {object} map[string]string
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var req AuthRequest
	c.BindJSON(&req)

	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid"})
		return
	}

	if !utils.CheckPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, _ := token.SignedString(config.GetJwtSecret())


	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

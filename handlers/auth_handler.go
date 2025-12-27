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
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password required"})
		return
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "hash failed"})
		return
	}

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
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var user models.User
	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if !utils.CheckPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	now := time.Now()

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"username": user.Username,
		"role": "user",
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"exp": now.Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(config.GetJwtSecret())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}


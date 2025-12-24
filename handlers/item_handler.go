package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-jwt-api/config"
	"go-jwt-api/models"
)

// @Security BearerAuth
// @Summary Get items
// @Tags Item
// @Produce json
// @Success 200 {array} models.ItemDTO
// @Router /items [get]
func GetItems(c *gin.Context) {
	var items []models.Item
	config.DB.Find(&items)

	var res []models.ItemDTO
	for _, item := range items {
		res = append(res, models.ItemDTO{
			ID:    item.ID,
			Name:  item.Name,
			Price: item.Price,
		})
	}

	c.JSON(http.StatusOK, res)
}


// @Security BearerAuth
// @Summary Create item
// @Tags Item
// @Accept json
// @Produce json
// @Param body body models.Item true "Item"
// @Success 200 {object} models.Item
// @Router /items [post]
func CreateItem(c *gin.Context) {
	var item models.Item
	c.BindJSON(&item)
	config.DB.Create(&item)
	c.JSON(http.StatusOK, item)
}

// @Security BearerAuth
// @Summary Update item
// @Tags Item
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param body body models.Item true "Item"
// @Success 200 {array} models.ItemDTO
// @Router /items/{id} [put]
func UpdateItem(c *gin.Context) {
	var item models.Item
	if err := config.DB.First(&item, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.BindJSON(&item)
	config.DB.Save(&item)
	c.JSON(http.StatusOK, item)
}

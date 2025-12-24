package seed

import (
	"fmt"

	"go-jwt-api/config"
	"go-jwt-api/models"
)

func SeedItems() {
	var count int64
	config.DB.Model(&models.Item{}).Count(&count)

	if count >= 100 {
		return
	}

	for i := 1; i <= 100; i++ {
		item := models.Item{
			Name:  fmt.Sprintf("Item %03d", i),
			Price: 1000 + (i * 10),
		}
		config.DB.Create(&item)
	}
}

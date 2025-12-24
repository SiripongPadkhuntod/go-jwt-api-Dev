package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name  string `gorm:"index" json:"name"`
	Price int    `gorm:"index" json:"price"`
}

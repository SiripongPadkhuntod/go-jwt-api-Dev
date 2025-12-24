package models

// import "gorm.io/gorm"

type Item struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Image string `json:"image"` // path หรือ url
}


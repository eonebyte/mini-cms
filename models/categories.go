package models

import (
	"myweb/config"
	// "myweb/models"
)

type Category struct {
	ID   int    `json:"id" form:"id" gorm:"primaryKey"`
	Name string `json:"name" form:"name"`
}

func GetAllCategories() ([]Category, error) {
	var categories []Category
	result := config.DB.Find(&categories)

	return categories, result.Error
}
func GetOneCategory(keywords string) (Category, error) {
	var categories Category
	result := config.DB.Where("name LIKE ?", "%"+keywords+"%").Preload("Articles").First(&categories)
	return categories, result.Error
}

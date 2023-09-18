package models

import (
	"myweb/config"
	// "myweb/models"
)

type Category struct {
	ID       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Articles []Article
}

func GetAllCategories() ([]Category, error) {
	var categories []Category
	err := config.DBGorm.Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func GetOneCategory(keywords string) (Category, error) {
	var categories Category
	result := config.DBGorm.Where("name LIKE ?", "%"+keywords+"%").Preload("Articles").First(&categories)
	return categories, result.Error
}

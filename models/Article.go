package models

import (
	"html/template"
	"myweb/config"
	"time"
)

// type Category struct {
//     ID int `json:"id" form:"id" gorm:"primaryKey"`
//     name string `json:"name" form:"name"`
//     Articles []Article
// }

type Article struct {
	ID          int           `json:"id" form:"id" gorm:"primary_key"`
	Title       string        `json:"title" form:"title"`
	Content     template.HTML `json:"content" form:"content"`
	Description string        `json:"description" form:"description"`
	CreatedAt   time.Time     `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" form:"updated_at"`
	CategoryId  string        `form:"category_id"`
}

func (article *Article) CreateArticle() error {
	err := config.DBGorm.Create(article).Error
	if err != nil {
		return err
	}

	return nil
}

func GetOneArticle(id int) (Article, error) {
	var article Article
	result := config.DBGorm.Where("id = ?", id).First(&article)
	return article, result.Error
}

func (article *Article) UpdateArticle(id int) error {
	if err := config.DBGorm.Model(&Article{}).Where("id = ?", id).Updates(article).Error; err != nil {
		return nil
	}
	return nil
}

func (article *Article) DeleteArticle() error {
	if err := config.DBGorm.Delete(article).Error; err != nil {
		return err
	}
	return nil
}

func GetAll(keywords string) ([]Article, error) {
	var article []Article
	result := config.DBGorm.Where("title LIKE ?", "%"+keywords+"%").Find(&article)

	return article, result.Error
}

func GetAllArticles() ([]Article, error) {
	var articles []Article
	err := config.DBGorm.Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

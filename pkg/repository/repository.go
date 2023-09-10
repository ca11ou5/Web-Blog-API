package repository

import (
	"blog/models"
	"gorm.io/gorm"
)

type Article interface {
	CreateArticle(article *models.Article) error
}

type Repository struct {
	Article
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Article: NewArticleRepository(db)}
}

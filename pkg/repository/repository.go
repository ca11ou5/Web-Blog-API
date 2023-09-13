package repository

import (
	"blog/models"
	"gorm.io/gorm"
)

type Article interface {
	GetAllArticles() *[]models.Article
	GetArticleByID(id uint) (*models.Article, error)
	CreateArticle(article *models.Article) (uint, error)
	DeleteArticleByID(id uint) error
}

type Authorization interface {
	CreateUser(user *models.User) (uint, error)
	GetUser(username, password string) (*models.User, error)
}

type Repository struct {
	Article
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Article:       NewArticleRepository(db),
		Authorization: NewAuthorizationRepository(db),
	}
}

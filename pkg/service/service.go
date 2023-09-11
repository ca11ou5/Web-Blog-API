package service

import (
	"blog/models"
	"blog/pkg/repository"
)

type Article interface {
	CreateArticle(article *models.Article) error
	GetArticleByID(id int) (*models.Article, error)
	GetAllArticles() (*[]models.Article, error)
	DeleteArticleByID(id int) error
}

type Service struct {
	Article
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Article: NewArticleService(repo)}
}

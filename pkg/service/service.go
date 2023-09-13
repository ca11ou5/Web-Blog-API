package service

import (
	"blog/models"
	"blog/pkg/repository"
)

type Article interface {
	GetAllArticles() *[]models.Article
	GetArticleByID(id uint) (*models.Article, error)
	CreateArticle(article *models.Article) (uint, error)
	DeleteArticleByID(id uint) error
}

type Authorization interface {
	CreateUser(user *models.User) (uint, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (uint, error)
}

type Service struct {
	Article
	Authorization
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Article:       NewArticleService(repo.Article),
		Authorization: NewAuthorizationService(repo.Authorization),
	}
}

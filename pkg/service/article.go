package service

import (
	"blog/models"
	"blog/pkg/repository"
)

type ArticleService struct {
	repo *repository.Repository
}

func NewArticleService(repo *repository.Repository) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) CreateArticle(article *models.Article) error {
	return s.repo.CreateArticle(article)
}

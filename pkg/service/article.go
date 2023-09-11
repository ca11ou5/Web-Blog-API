package service

import (
	"blog/models"
	"blog/pkg/repository"
	"github.com/gosimple/slug"
)

type ArticleService struct {
	repo *repository.Repository
}

func NewArticleService(repo *repository.Repository) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) CreateArticle(article *models.Article) error {
	article.Slug = slug.Make(article.Name)
	return s.repo.CreateArticle(article)
}

func (s *ArticleService) GetArticleByID(id int) (*models.Article, error) {
	return s.repo.GetArticleByID(id)
}

func (s *ArticleService) GetAllArticles() (*[]models.Article, error) {
	return s.repo.GetAllArticles()
}

func (s *ArticleService) DeleteArticleByID(id int) error {
	return s.repo.DeleteArticleByID(id)
}

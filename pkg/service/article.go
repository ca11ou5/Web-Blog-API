package service

import (
	"blog/models"
	"blog/pkg/repository"
	"github.com/gosimple/slug"
)

type ArticleService struct {
	repo repository.Article
}

func NewArticleService(repo repository.Article) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) GetAllArticles() *[]models.Article {
	return s.repo.GetAllArticles()
}

func (s *ArticleService) GetArticleByID(id uint) (*models.Article, error) {
	return s.repo.GetArticleByID(id)
}

func (s *ArticleService) CreateArticle(article *models.Article) (uint, error) {
	article.Slug = slug.Make(article.Name)
	return s.repo.CreateArticle(article)
}

func (s *ArticleService) DeleteArticleByID(id uint) error {
	return s.repo.DeleteArticleByID(id)
}

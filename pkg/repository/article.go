package repository

import (
	"blog/models"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db: db}
}

func (r *ArticleRepository) GetAllArticles() *[]models.Article {
	var articles []models.Article
	r.db.Preload("Texts").Preload("Codes").Preload("Titles").Preload("Images").Find(&articles)
	return &articles
}

func (r *ArticleRepository) GetArticleByID(id uint) (*models.Article, error) {
	article := models.Article{}
	err := r.db.Take(&article, id)
	return &article, err.Error
}

func (r *ArticleRepository) CreateArticle(article *models.Article) (uint, error) {
	err := r.db.Create(&article)
	return article.ID, err.Error
}

func (r *ArticleRepository) DeleteArticleByID(id uint) error {
	article := models.Article{}
	err := r.db.Preload("Texts").Preload("Codes").Preload("Titles").Preload("Images").First(&article, id).Delete(&article)
	return err.Error
}

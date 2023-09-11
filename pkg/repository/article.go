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

func (r *ArticleRepository) CreateArticle(article *models.Article) error {
	r.db.Create(&article)
	return nil
}

func (r *ArticleRepository) GetArticleByID(id int) (*models.Article, error) {
	article := models.Article{}
	r.db.Take(&article, id)
	return &article, nil
}

func (r *ArticleRepository) GetAllArticles() (*[]models.Article, error) {
	var articles []models.Article
	r.db.Preload("Texts").Preload("Codes").Preload("Titles").Preload("Images").Find(&articles)
	return &articles, nil
}

func (r *ArticleRepository) DeleteArticleByID(id int) error {
	article := models.Article{}
	r.db.Preload("Texts").Preload("Codes").Preload("Titles").Preload("Images").First(&article, id)
	r.db.Delete(&article)
	return nil
}

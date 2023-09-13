package repository

import (
	"blog/models"
	"gorm.io/gorm"
)

type AuthorizationRepository struct {
	db *gorm.DB
}

func NewAuthorizationRepository(db *gorm.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db: db}
}

func (r *AuthorizationRepository) CreateUser(user *models.User) (uint, error) {
	err := r.db.Create(&user)
	return user.ID, err.Error
}

func (r *AuthorizationRepository) GetUser(username string, password string) (*models.User, error) {
	user := models.User{}
	err := r.db.Where("username = ? AND password = ?", username, password).First(&user)

	return &user, err.Error
}

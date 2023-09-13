package service

import (
	"blog/models"
	"blog/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthorizationService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService) CreateUser(user *models.User) (uint, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

const salt = "d927js213jsa4k2a1o"
const signingKey = "mdsa213kfos@dsm(*asz"

type tokenClaims struct {
	jwt.StandardClaims
	UserID uint
}

func (s *AuthorizationService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		},
		UserID: user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthorizationService) ParseToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type")
	}

	return claims.UserID, nil
}

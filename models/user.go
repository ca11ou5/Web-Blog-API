package models

type User struct {
	ID       uint   `json:"-"`
	Name     string `gorm:"not null" binding:"required"`
	Username string `gorm:"not null;unique" binding:"required"`
	Password string `gorm:"not null" binding:"required"`
}

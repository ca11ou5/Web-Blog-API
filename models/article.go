package models

import "time"

type Article struct {
	ID        uint      `json:"-"`
	Name      string    `gorm:"not null" binding:"required"`
	Slug      string    `gorm:"unique_index" json:"-"`
	Status    bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Images    []Image   `gorm:"constraint:OnDelete:CASCADE;"`
	Texts     []Text    `gorm:"constraint:OnDelete:CASCADE;"`
	Codes     []Code    `gorm:"constraint:OnDelete:CASCADE;"`
	Titles    []Title   `gorm:"constraint:OnDelete:CASCADE;"`
}

type Image struct {
	ID        uint `json:"-"`
	ArticleID uint `json:"-"`
	Order     uint
	Content   []byte
}

type Text struct {
	ID        uint `json:"-"`
	ArticleID uint `json:"-"`
	Order     uint
	Content   string
}

type Code struct {
	ID        uint `json:"-"`
	ArticleID uint `json:"-"`
	Order     uint
	Content   string
}

type Title struct {
	ID        uint `json:"-"`
	ArticleID uint `json:"-"`
	Order     uint
	Content   string
}

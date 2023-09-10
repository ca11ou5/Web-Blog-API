package models

import "time"

type Article struct {
	ID        uint
	Name      string
	Status    bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Images    []Image
	Texts     []Text
	Codes     []Code
	Titles    []Title
}

type Image struct {
	ID        uint
	ArticleID uint
	Order     uint
	Content   []byte
}

type Text struct {
	ID        uint
	ArticleID uint
	Order     uint
	Content   string
}

type Code struct {
	ID        uint
	ArticleID uint
	Order     uint
	Content   string
}

type Title struct {
	ID        uint
	ArticleID uint
	Order     uint
	Content   string
}

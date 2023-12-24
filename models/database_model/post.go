package database_model

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Content     string `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
	User        User   `gorm:"foreignKey:UserID"`
	NumLikes    uint   `gorm:"default:0"`
	NumDislikes uint   `gorm:"default:0"`
	NumReports  uint   `gorm:"default:0"`
	NumShares   uint   `gorm:"default:0"`
	ImageURL    string
}

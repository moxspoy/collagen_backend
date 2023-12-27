package database_model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content     string `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
	User        User   `gorm:"foreignKey:UserID"`
	PostID      uint   `gorm:"not null"`
	Post        Post   `gorm:"foreignKey:PostID"`
	NumLikes    uint   `gorm:"default:0"`
	NumDislikes uint   `gorm:"default:0"`
	NumReports  uint   `gorm:"default:0"`
}

package database_model

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Title       string    `gorm:"not null"`
	Content     string    `gorm:"not null"`
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
	NumLikes    uint      `gorm:"default:0"`
	NumDislikes uint      `gorm:"default:0"`
	NumReports  uint      `gorm:"default:0"`
	Comments    []Comment
}

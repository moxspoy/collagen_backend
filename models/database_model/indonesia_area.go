package database_model

import (
	"gorm.io/gorm"
)

type IndonesiaArea struct {
	gorm.Model
	Code string `binding:"required"`
	Name string `binding:"required"`
}

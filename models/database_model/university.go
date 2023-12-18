package database_model

import (
	"gorm.io/gorm"
)

type University struct {
	gorm.Model
	Name    string `binding:"required"`
	NameEn  string
	Acronym string

	// Active (1) or Not Active (2)
	Status int16

	// State owned (1) or private (2)
	Category int16

	Website         string
	YearFounded     int16
	IndonesiaAreaID uint
	IndonesiaArea   IndonesiaArea `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (University) TableName() string {
	return "university"
}

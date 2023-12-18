package database_model

import (
	"gorm.io/gorm"
)

// database from https://github.com/cahyadsn/wilayah/blob/master/db/wilayah.sql

type IndonesiaArea struct {
	gorm.Model
	Code string `binding:"required"`
	Name string `binding:"required"`
}

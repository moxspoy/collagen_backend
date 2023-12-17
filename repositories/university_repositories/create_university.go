package university_repositories

import (
	"collagen/config/database"
	"collagen/models/database_model"
	"gorm.io/gorm"
)

func CreateUniversity(data database_model.University) *gorm.DB {
	return database.DB.Create(&data)
}

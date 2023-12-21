package university_repositories

import (
	"collagen/config/database"
	"collagen/models/database_model"
)

func GetAllUniversities() []database_model.University {
	var universities []database_model.University
	database.DB.Find(&universities)
	return universities
}

package user_repository

import (
	"collagen/config/database"
	"collagen/models/database_model"
)

func GetOneUserById(userId uint) database_model.User {
	var user database_model.User
	database.DB.Where("id = ?", userId).First(&user)
	return user
}

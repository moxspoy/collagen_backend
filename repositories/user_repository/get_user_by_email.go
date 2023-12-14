package user_repository

import (
	"collagen/config/database"
	"collagen/models/database_model"
)

func GetOneUserByEmail(email string) database_model.User {
	var user database_model.User
	database.DB.Where("email = ?", email).First(&user)
	return user
}

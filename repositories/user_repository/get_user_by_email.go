package user_repository

import (
	"flop/config/database"
	"flop/models/database_model"
)

func GetOneUserByEmail(email string) database_model.User {
	var user database_model.User
	database.DB.Where("email = ?", email).First(&user)
	return user
}

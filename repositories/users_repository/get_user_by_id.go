package users_repository

import (
	"flop/config/database"
	"flop/models/database_model"
)

func GetOneUserById(userId uint) database_model.Users {
	var user database_model.Users
	database.DB.Where("id = ?", userId).First(&user)
	return user
}

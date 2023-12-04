package users

import (
	"flop/config/database"
	"flop/models/database_model"
)

func GetOneUserByEmail(email string) database_model.Users {
	var user database_model.Users
	database.DB.Where("email = ?", email).First(&user)
	return user
}

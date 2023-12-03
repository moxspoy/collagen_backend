package users

import (
	"flop/config/database"
	"flop/models"
)

func GetOneUserByEmail(email string) models.Users {
	var user models.Users
	database.DB.Where("email = ?", email).First(&user)
	return user
}

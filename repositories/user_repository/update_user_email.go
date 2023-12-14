package user_repository

import (
	"collagen/config/database"
	"collagen/models/database_model"
)

func UpdateUserEmail(user *database_model.User, userId uint, newEmail string) {
	database.DB.Model(&user).Where("id = ?", userId).UpdateColumn("email", newEmail)
	newUser := GetOneUserByEmail(newEmail)
	user.ID = newUser.ID
}

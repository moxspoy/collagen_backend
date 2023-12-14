package user_repository

import (
	"collagen/config/database"
	"collagen/models/database_model"
)

func UpdateUserEmail(user *database_model.User, oldEmail string, newEmail string) {
	database.DB.Model(&user).Where("email = ?", oldEmail).UpdateColumn("email", newEmail)
	newUser := GetOneUserByEmail(newEmail)
	user.ID = newUser.ID
}

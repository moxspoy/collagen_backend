package user_repository

import (
	"flop/config/database"
	"flop/models/database_model"
)

func UpdateUserEmail(user *database_model.User, oldEmail string, newEmail string) {
	database.DB.Model(&user).Where("email = ?", oldEmail).UpdateColumn("email", newEmail)
	newUser := GetOneUserByEmail(newEmail)
	user.ID = newUser.ID
}

package users

import (
	"flop/config/database"
	"flop/models/database_model"
)

func UpdateUserEmail(user *database_model.Users, oldEmail string, newEmail string) {
	database.DB.Model(&user).Where("email = ?", oldEmail).UpdateColumn("email", newEmail)
	newUser := GetOneUserByEmail(newEmail)
	user.Id = newUser.Id
}

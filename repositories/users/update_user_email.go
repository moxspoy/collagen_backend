package users

import (
	"flop/config/database"
	"flop/models"
)

func UpdateUserEmail(oldEmail string, newEmail string) {
	database.DB.Model(&models.Users{}).Where("email = ?", oldEmail).UpdateColumn("email", newEmail)
}

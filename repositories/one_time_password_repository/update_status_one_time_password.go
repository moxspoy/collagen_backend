package one_time_password_repository

import (
	"flop/config/database"
	"flop/models/database_model"
)

func UpdateStatusOneTimePassword(id uint, status int) error {
	var otpFromDatabase database_model.OneTimePassword
	dbc := database.DB.Model(&otpFromDatabase).Where("id = ?", id).UpdateColumn("status", status)
	if dbc.Error != nil {
		return dbc.Error
	}

	return nil
}

func UpdateStatusOneTimePasswordToUsed(id uint) error {
	return UpdateStatusOneTimePassword(id, 1)
}

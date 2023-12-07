package one_time_password_repository

import (
	"flop/config/database"
	"flop/models/database_model"
)

func UpdateStatusOneTimePassword(id uint, status int) (database_model.OneTimePassword, error) {
	var otpFromDatabase database_model.OneTimePassword
	dbc := database.DB.Model(&otpFromDatabase).Where("id = ?", id).UpdateColumn("status", status)
	if dbc.Error != nil {
		return otpFromDatabase, dbc.Error
	}

	return otpFromDatabase, nil
}

func UpdateStatusOneTimePasswordToUsed(id uint) (database_model.OneTimePassword, error) {
	return UpdateStatusOneTimePassword(id, 1)
}

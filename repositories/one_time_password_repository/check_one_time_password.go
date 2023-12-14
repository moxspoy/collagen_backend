package one_time_password_repository

import (
	"collagen/config/database"
	"collagen/helper/security_helper"
	"collagen/models/database_model"
	"errors"
)

func CheckOneTimePassword(userId uint, otpRequest string) error {
	var otpFromDatabase database_model.OneTimePassword
	tx := database.DB.Last(&otpFromDatabase, "user_id = ? and status = ? and token = ?", userId, 0, otpRequest)
	if tx.Error != nil {
		return tx.Error
	}
	if security_helper.IsValidOtp(otpFromDatabase) {
		err := UpdateStatusOneTimePasswordToUsed(otpFromDatabase.ID)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("OTP is not valid")
}

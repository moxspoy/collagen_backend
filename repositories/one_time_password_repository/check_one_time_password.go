package one_time_password_repository

import (
	"errors"
	"flop/config/database"
	"flop/helper/security_helper"
	"flop/models/database_model"
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

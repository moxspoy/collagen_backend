package one_time_password_repository

import (
	"errors"
	"flop/config/database"
	"flop/helper/security_helper"
	"flop/models/database_model"
)

func CheckOneTimePassword(userId uint, otpRequest string) (database_model.OneTimePassword, error) {
	var otpFromDatabase database_model.OneTimePassword
	dbc := database.DB.Last(&otpFromDatabase, "user_id = ? and status = ? and token = ?", userId, 0, otpRequest)
	if dbc.Error != nil {
		return otpFromDatabase, dbc.Error
	}
	if security_helper.IsValidOtp(otpFromDatabase) {
		return otpFromDatabase, nil
	}
	return otpFromDatabase, errors.New("OTP is not valid")
}

package security_helper

import (
	"collagen/models/database_model"
	"crypto/rand"
	"time"
)

const otpChars = "1234567890"

func GenerateOTP() (string, error) {
	length := 6
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func IsValidOtp(otpFromDatabase database_model.OneTimePassword) bool {
	expiredAt := otpFromDatabase.ExpiredAt
	isExpired := time.Now().Unix() > expiredAt.Unix()
	if isExpired {
		return false
	}
	if otpFromDatabase.Status != 0 {
		return false
	}
	return true
}

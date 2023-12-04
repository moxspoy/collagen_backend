package users

import (
	"flop/config/database"
	database2 "flop/models/database_model"
)

func GetOneUserById(userId uint) database2.Users {
	var user database2.Users
	database.DB.Where("id = ?", userId).First(&user)
	return user
}

package upload_helper

import (
	"collagen/helper/api_response_helper"
	"collagen/middleware"
	"collagen/models/database_model"
	"collagen/repositories/user_detail_repository"
	"collagen/repositories/user_repository"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func UploadImageForKyc(c *gin.Context, path string) {
	userId := middleware.GetUserIdFromJWT(c)
	file, err := c.FormFile("file")
	if err != nil {
		api_response_helper.GenerateErrorResponse(c, err)
		return
	}
	if file.Filename == "" {
		api_response_helper.GenerateErrorResponse(c, errors.New("file must not be empty"))
		return
	}
	dest := "./assets/kyc/" + path + strconv.FormatUint(uint64(userId), 10) + ".png"

	err = c.SaveUploadedFile(file, dest)
	if err != nil {
		api_response_helper.GenerateErrorResponse(c, err)
		return
	}

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	var userDetail database_model.UserDetail
	userDetail.UserID = userId
	if strings.Contains(path, "identity_and_selfie") {
		userDetail.IdentityAndSelfieImageURL = dest
	} else if strings.Contains(path, "identity") {
		userDetail.IdentityImageURL = dest
	} else {
		userDetail.SelfieImageURL = dest
	}
	trx := user_detail_repository.UpdateUserDetail(&userDetail)

	user := user_repository.GetOneUserById(userId)
	userDetail.User = user

	if trx.Error != nil {
		api_response_helper.GenerateErrorResponse(c, trx.Error)
		return
	}
	api_response_helper.GenerateSuccessResponse(c, "Update user's detail successful", userDetail)
}

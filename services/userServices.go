package services

import (
	"codemax_assignment/helpers/loggerhelper"
	"codemax_assignment/helpers/tokenhelper"
	"codemax_assignment/models"
)

// GenerateApiToken- generate token as per username and password
func GenerateApiToken(user models.User) (string, error) {
	if setUserErr := models.SetUser(user.UserName, user.Password); setUserErr != nil {
		loggerhelper.LogError("FAILED_TO_ASSIGN_USER_TOKEN", setUserErr)
		return "", setUserErr
	}
	token, genTokenErr := tokenhelper.GenerateToken(user.UserName, user.Password)
	if genTokenErr != nil {
		loggerhelper.LogError("FAILED_TO_GENERATE_TOKEN::", genTokenErr)
		return "", genTokenErr
	}
	return token, nil
}

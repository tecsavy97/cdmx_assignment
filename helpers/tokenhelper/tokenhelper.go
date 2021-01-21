package tokenhelper

import (
	"codemax_assignment/models"
	b64 "encoding/base64"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

//AccessByToken - get token from header and decode it
func AccessByToken(c *gin.Context) error {
	if decodeErr := DecodeToken(c.GetHeader("Authorization")); decodeErr != nil {
		return decodeErr
	}
	return nil
}

//DecodeToken - decode token and access the route
func DecodeToken(token string) error {
	tokenDec, decErr := b64.StdEncoding.DecodeString(token)
	if decErr != nil {
		return decErr
	}
	userDetails := strings.Split(string(tokenDec), ":")
	if len(userDetails) <= 0 {
		return errors.New("TOKEN_INVALID")
	}
	userName, userPassword := userDetails[0], userDetails[1]
	user, getUserErr := models.GetUser(userName)
	if getUserErr != nil {
		return getUserErr
	}
	if userName != user.UserName && userPassword != user.Password {
		return errors.New("FAILED_TO_AUTHORIZE_API_TOKEN")
	}
	return nil
}

// GenerateToken - helper to generate token as per given data
func GenerateToken(userName, userPassword string) (string, error) {
	if userName == "" && userPassword == "" {
		return "", errors.New("FAILED_TO_GENERATE_API_TOKEN")
	}
	token := b64.StdEncoding.EncodeToString([]byte(userName + ":" + userPassword))
	return token, nil
}

package handlers

import (
	"codemax_assignment/helpers/loggerhelper"
	"codemax_assignment/models"
	"codemax_assignment/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GenerateToken -  handler to handle request body and return a token
func GenerateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := models.NewUser()
		if bindErr := c.Bind(&user); bindErr != nil {
			loggerhelper.LogError("FAILED_TO_BIND_REQUEST_BODY::", bindErr)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": bindErr.Error(), "message": "FAILED_TO_BIND_REQUEST_BODY"})
			return
		}
		token, genTokenErr := services.GenerateApiToken(user)
		if genTokenErr != nil {
			loggerhelper.LogError("FAILED_TO_GENERATE_API_TOKEN::", genTokenErr)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": genTokenErr.Error(), "message": "FAILED_TO_GENERATE_API_TOKEN"})
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Use the Tokena and set it in 'Authorization' Header to send email via /send-email route", "token": token})
	}
}

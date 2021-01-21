package handlers

import (
	"codemax_assignment/helpers/loggerhelper"
	"codemax_assignment/models"
	"codemax_assignment/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//SendEmailHandler - handler to handle request body data and send email as per given request
func SendEmailHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := models.NewEmail()
		if bindErr := c.Bind(&email); bindErr != nil {
			loggerhelper.LogError("FAILED_TO_BIND_REQUEST_BODY::", bindErr)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": bindErr.Error(), "message": "FAILED_TO_BIND_REQUEST_BODY"})
			return
		}
		if sendErr := services.SendEmail(email); sendErr != nil {
			loggerhelper.LogError("FAILED_TO_SEND_EMAIL::", sendErr)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": sendErr.Error(), "message": "FAILED_TO_SEND_EMAIL"})
			return
		}
		receipients := email.To
		receipients = append(receipients, email.CC...)
		receipients = append(receipients, email.BCC...)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"message": "Email sent to " + strings.Join(receipients, ",")})
	}
}

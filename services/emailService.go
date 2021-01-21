package services

import (
	"codemax_assignment/helpers/emailhelper"
	"codemax_assignment/helpers/loggerhelper"
	"codemax_assignment/models"
	"io/ioutil"
)

// InitConfig - initialise emailconfig for single email instance
func InitConfig() error {
	fileData, readErr := ioutil.ReadFile(models.AppConfig.EmailConfPath)
	if readErr != nil {
		loggerhelper.LogError("FAILED_TO_READ_JSON_FILE::", readErr)
		return readErr
	}
	if initErr := emailhelper.InitConfig(fileData); initErr != nil {
		loggerhelper.LogError("FAILED_TO_INIT_EMAIL_CONFIG::", initErr)
	}
	return nil
}

// SendEmail - service to send email to multiple users
func SendEmail(email models.Email) error {
	if initErr := InitConfig(); initErr != nil {
		loggerhelper.LogError("FAILED_TO_INIT_EMAIL_CONF::", initErr)
		return initErr
	}
	em := emailhelper.NewMail(email.To, email.CC, email.BCC, email.From, email.ReplyTo, email.Subject, email.Body)
	if sendErr := em.SendMail(); sendErr != nil {
		loggerhelper.LogError("FAILED_TO_SEND_EMAIL::", sendErr)
		return sendErr
	}
	return nil
}

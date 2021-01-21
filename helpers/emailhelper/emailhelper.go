package emailhelper

import (
	"codemax_assignment/models"
	"encoding/json"
	"errors"
	"net/smtp"
	"sync"

	"gopkg.in/gomail.v2"
)

var (
	// map of emailConfig instances
	instances map[string]EmailConfig
	// to initialise map of emailConfig once
	onceMulti sync.Once
	//default host if any host is set to be default in th struct
	defaultHost string
	//single emailConfig instance
	emailConfig EmailConfig
	// to initialise emailconfig once
	once sync.Once
)

// EmailConfig - EmailConfig body
type EmailConfig struct {
	Auth      Auth   `json:"auth_keys"`
	HostName  string `json:"hostName"`
	Server    string `json:"server"`
	Port      int    `json:"port"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	IsDefault bool   `json:"isDefault"`
}

// Auth - auth from smtp auth keys
type Auth struct {
	Identity  string `json:"identity"`
	PublicKey string `json:"publicKey"`
	SecretKey string `json:"secretKey"`
}

// Email struct for helper
type Email struct {
	from    string
	replyTo string
	to      []string
	cc      []string
	bcc     []string
	subject string
	body    string
}

// InitUsingJSON initializes multiples Email Connections for given JSON data
func InitUsingJSON(configs []EmailConfig) {
	onceMulti.Do(func() {
		instances = make(map[string]EmailConfig)
		for _, config := range configs {
			instances[config.HostName] = config
			if config.IsDefault {
				defaultHost = config.HostName
			}
		}
	})
}

// initialise emailconfig for single instance without any map
func InitConfig(data []byte) error {
	var initConfErr error
	once.Do(func() {
		err := json.Unmarshal([]byte(data), &emailConfig)
		initConfErr = err
		emailConfig.Auth.PublicKey = models.AppConfig.PublicKey
		emailConfig.Auth.SecretKey = models.AppConfig.SecretKey
	})
	return initConfErr
}

//SendMail - send email for single email instance
func (email *Email) SendMail() error {
	message := gomail.NewMessage()
	message.SetHeader("From", email.from)
	message.SetHeader("To", email.to...)
	message.SetHeader("Reply-To", email.replyTo)
	message.SetHeader("Cc", email.cc...)
	message.SetHeader("Bcc", email.bcc...)
	message.SetHeader("Subject", email.subject)
	message.SetBody("text/plain", email.body)
	dailer := gomail.Dialer{Host: emailConfig.Server, Port: emailConfig.Port, Username: emailConfig.Username, Password: emailConfig.Password, Auth: smtp.PlainAuth(emailConfig.Auth.Identity, emailConfig.Auth.PublicKey, emailConfig.Auth.SecretKey, emailConfig.Server)}
	if err := dailer.DialAndSend(message); err != nil {
		return err
	}
	return nil
}

// SendMailByHost - To get EmailConfig by hostname and then sending the email as per the Email Body
func (email *Email) SendMailByHost(hostName string) error {
	config := EmailConfig{}
	if hostName == "" {
		tmp, ok := instances[defaultHost]
		if !ok {
			return errors.New("Host not found :" + hostName)
		}
		config = tmp
	} else {
		tmp, ok := instances[hostName]
		if !ok {
			return errors.New("Host not found :" + hostName)
		}
		config = tmp
	}
	message := gomail.NewMessage()
	message.SetHeader("From", email.from)
	message.SetHeader("To", email.to...)
	message.SetHeader("Reply-To", email.replyTo)
	message.SetHeader("Cc", email.cc...)
	message.SetHeader("Bcc", email.bcc...)
	message.SetHeader("Subject", email.subject)
	message.SetBody("text/plain", email.body)
	dailer := gomail.Dialer{Host: config.Server, Port: config.Port, Username: config.Username, Password: config.Password, Auth: smtp.PlainAuth(config.Auth.Identity, config.Auth.PublicKey, config.Auth.SecretKey, config.Server)}
	if err := dailer.DialAndSend(message); err != nil {
		return err
	}
	return nil
}

// new Email body to be use to send emails
func NewMail(to, cc, bcc []string, from, replyTo, subject, body string) *Email {
	return &Email{
		to:      to,
		cc:      cc,
		bcc:     bcc,
		from:    from,
		replyTo: replyTo,
		subject: subject,
		body:    body,
	}
}

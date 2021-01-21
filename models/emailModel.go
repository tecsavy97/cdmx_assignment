package models

// email body to be send in request
type Email struct {
	To      []string `json:"to"`
	CC      []string `json:"cc,omitempty"`
	BCC     []string `json:"bcc,omitempty"`
	From    string   `json:"from"`
	ReplyTo string   `json:"replyTo,omitempty"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

//NewEmail - to bind body of requesr
func NewEmail() Email {
	return Email{}
}

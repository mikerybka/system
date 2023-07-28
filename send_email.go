package system

import (
	"encoding/json"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var emailConfig struct {
	From struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"from"`
	Sendgrid struct {
		APIKey string `json:"api_key"`
	} `json:"sendgrid"`
}

func init() {
	b, _ := os.ReadFile("/etc/email.json")
	json.Unmarshal(b, &emailConfig)
}

func SendEmail(toAddress, toName, subject, textMessage, htmlMessage string) error {
	from := mail.NewEmail(emailConfig.From.Name, emailConfig.From.Email)
	to := mail.NewEmail(toName, toAddress)
	email := mail.NewSingleEmail(from, subject, to, textMessage, htmlMessage)
	client := sendgrid.NewSendClient(emailConfig.Sendgrid.APIKey)
	_, err := client.Send(email)
	return err
}

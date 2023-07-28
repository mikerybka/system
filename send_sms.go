package system

import (
	"encoding/json"
	"os"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)

var smsConfig struct {
	Twilio struct {
		AccountSID string `json:"account_sid"`
		AuthToken  string `json:"auth_token"`
	} `json:"twilio"`
	From string `json:"from"`
}

func init() {
	b, _ := os.ReadFile("/etc/sms.json")
	json.Unmarshal(b, &smsConfig)
	os.Setenv("TWILIO_ACCOUNT_SID", smsConfig.Twilio.AccountSID)
	os.Setenv("TWILIO_AUTH_TOKEN", smsConfig.Twilio.AuthToken)
}

func SendSMS(to, message string) error {
	client := twilio.NewRestClient()
	params := &api.CreateMessageParams{}
	params.SetFrom(smsConfig.From)
	params.SetTo(to)
	params.SetBody(message)
	_, err := client.Api.CreateMessage(params)
	return err
}

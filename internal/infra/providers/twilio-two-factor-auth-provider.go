package providers

import (
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioTwoFactorAuthProvider struct {
	Client *twilio.RestClient
}

func (tp *TwilioTwoFactorAuthProvider) Send(from, to, message string) error {
	params := &openapi.CreateMessageParams{}
	params.SetFrom(from)
	params.SetTo(to)
	params.SetBody(message)

	_, err := tp.Client.Api.CreateMessage(params)
	if err != nil {
		return err
	}

	return nil
}

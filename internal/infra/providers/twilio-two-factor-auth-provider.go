package providers

import (
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioTwoFactorAuthProvider struct {
	Client *twilio.RestClient
}

func (tp *TwilioTwoFactorAuthProvider) Send(from, to, message string) errors.IAppError {
	params := &openapi.CreateMessageParams{}
	params.SetFrom(from)
	params.SetTo(to)
	params.SetBody(message)

	_, err := tp.Client.Api.CreateMessage(params)
	if err != nil {
		logger.Error("Error trying to send sms with twilio.", err)
		return errors.NewAppError("internal server error.", 500)
	}

	return nil
}

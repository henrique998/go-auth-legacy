package providers

import "github.com/resend/resend-go/v2"

type ResendEmailProvider struct {
	ApiKey string
}

func (ep *ResendEmailProvider) SendMail(to string, subject string, body string) error {
	client := resend.NewClient(ep.ApiKey)

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{to},
		Subject: subject,
		Text:    body,
		Cc:      nil,
		Bcc:     nil,
		ReplyTo: "henriquemonteiro037@gmail.com",
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		return err
	}

	return nil
}

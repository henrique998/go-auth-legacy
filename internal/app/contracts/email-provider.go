package contracts

type EmailProvider interface {
	SendMail(to string, subject string, body string) error
}

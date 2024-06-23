package contracts

type TwoFactorAuthProvider interface {
	Send(from, to, message string) error
}

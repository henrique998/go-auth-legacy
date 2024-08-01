package contracts

type AuthSocialProvider interface {
	GetGoogleAccessToken(code string) (string, error)
}

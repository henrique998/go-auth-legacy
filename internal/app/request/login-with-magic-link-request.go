package request

type LoginWithMagicLinkRequest struct {
	Code      string
	IP        string
	UserAgent string
}

package request

type LoginWithGoogleRequest struct {
	Code      string `json:"code"`
	IP        string
	UserAgent string
}

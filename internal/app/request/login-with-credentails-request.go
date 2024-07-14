package request

type LoginWithCredentialsRequest struct {
	Email     string `json:"email"`
	Pass      string `json:"password"`
	IP        string
	UserAgent string
}

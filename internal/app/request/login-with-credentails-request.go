package request

type LoginWithCredentialsRequest struct {
	Email      string `json:"email"`
	Pass       string `json:"password"`
	IP         string
	DeviceName string `json:"deviceName"`
	UserAgent  string `json:"userAgent`
	Platform   string `json:"platform"`
}

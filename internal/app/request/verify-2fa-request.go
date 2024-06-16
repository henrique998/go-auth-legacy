package request

type Verify2faRequest struct {
	Code      string `json:"code"`
	AccountId string
}

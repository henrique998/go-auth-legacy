package request

type NewPassRequest struct {
	Code                string `json:"code"`
	NewPass             string `json:"newPass"`
	NewPassConfirmation string `json:"newPassConfirmation"`
}

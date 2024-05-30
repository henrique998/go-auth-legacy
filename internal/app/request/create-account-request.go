package request

type CreateAccountRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Pass  string `json:"password"`
	Phone string `json:"phone"`
}

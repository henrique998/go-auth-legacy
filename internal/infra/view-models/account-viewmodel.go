package viewmodels

import (
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/response"
)

func ToAccountHttp(account entities.Account) response.AccountResponse {
	return response.AccountResponse{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		Phone:     account.Phone,
		CreatedAt: account.CreatedAt,
	}
}

package viewmodels

import (
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/response"
)

func ToUserHttp(user entities.IAccount) response.UserResponse {
	return response.UserResponse{
		ID:        user.GetId(),
		Name:      user.GetName(),
		Email:     user.GetEmail(),
		Phone:     user.GetPhone(),
		CreatedAt: user.GetCreatedAt(),
	}
}

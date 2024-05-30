package mappers

import (
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
)

type DbUserData struct {
	ID               string
	Name             string
	Email            string
	Pass             string
	Phone            string
	Is2faEnabled     bool
	LastLogin        *time.Time
	LastLoginIp      *string
	LastLoginCountry *string
	LastLoginCity    *string
	CreatedAt        time.Time
	UpdatedAt        *time.Time
}

func MapToUser(data DbUserData) entities.IUser {
	var lastLogin time.Time
	if data.LastLogin != nil {
		lastLogin = *data.LastLogin
	}

	var lastLoginIp string
	if data.LastLoginIp != nil {
		lastLoginIp = *data.LastLoginIp
	}

	var lastLoginCountry string
	if data.LastLoginCountry != nil {
		lastLoginCountry = *data.LastLoginCountry
	}

	var lastLoginCity string
	if data.LastLoginCity != nil {
		lastLoginCity = *data.LastLoginCity
	}

	var updatedAt time.Time
	if data.UpdatedAt != nil {
		updatedAt = *data.UpdatedAt
	}

	return entities.NewExistingUser(
		data.ID,
		data.Name,
		data.Email,
		data.Pass,
		data.Phone,
		data.Is2faEnabled,
		lastLogin,
		lastLoginIp,
		lastLoginCountry,
		lastLoginCity,
		data.CreatedAt,
		updatedAt,
	)
}

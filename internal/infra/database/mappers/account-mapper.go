package mappers

import (
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
)

type DbAccountData struct {
	ID               string
	Name             string
	Email            string
	Pass             *string
	Phone            *string
	ProviderId       *string
	Is2faEnabled     bool
	IsEmailVerified  bool
	LastLoginAt      *time.Time
	LastLoginIp      *string
	LastLoginCountry *string
	LastLoginCity    *string
	CreatedAt        time.Time
	UpdatedAt        *time.Time
}

func MapToAccount(data DbAccountData) entities.Account {
	var Pass string
	if data.Pass != nil {
		Pass = *data.Pass
	}

	var Phone string
	if data.Phone != nil {
		Phone = *data.Phone
	}

	var providerId string
	if data.ProviderId != nil {
		providerId = *data.ProviderId
	}

	var lastLoginAt time.Time
	if data.LastLoginAt != nil {
		lastLoginAt = *data.LastLoginAt
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

	return *entities.NewExistingAccount(
		data.ID,
		data.Name,
		data.Email,
		Pass,
		Phone,
		providerId,
		data.Is2faEnabled,
		data.IsEmailVerified,
		lastLoginAt,
		lastLoginIp,
		lastLoginCountry,
		lastLoginCity,
		data.CreatedAt,
		updatedAt,
	)
}

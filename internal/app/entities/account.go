package entities

import (
	"time"

	"github.com/henrique998/go-auth/internal/infra/utils"
)

type Account struct {
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

func NewAccount(name, email, pass, phone, providerId string) *Account {
	var passVal, phoneVal, providerIdVal *string

	if pass == "" {
		passVal = nil
	} else {
		passVal = &pass
	}

	if phone == "" {
		phoneVal = nil
	} else {
		phoneVal = &phone
	}

	if providerId == "" {
		providerIdVal = nil
	} else {
		providerIdVal = &providerId
	}

	return &Account{
		ID:               utils.GenerateUUID(),
		Name:             name,
		Email:            email,
		Pass:             passVal,
		Phone:            phoneVal,
		ProviderId:       providerIdVal,
		Is2faEnabled:     false,
		IsEmailVerified:  false,
		LastLoginAt:      nil,
		LastLoginIp:      nil,
		LastLoginCountry: nil,
		LastLoginCity:    nil,
		CreatedAt:        time.Now(),
		UpdatedAt:        nil,
	}
}

func NewExistingAccount(id, name, email, pass, phone, providerId string, is2faEnabled, isEmailVerified bool, lastLogin time.Time, lastLoginIp, lastLoginCountry, lastLoginCity string, createdAt time.Time, updatedAt time.Time) *Account {
	return &Account{
		ID:               id,
		Name:             name,
		Email:            email,
		Pass:             &pass,
		Phone:            &phone,
		ProviderId:       &providerId,
		Is2faEnabled:     is2faEnabled,
		IsEmailVerified:  isEmailVerified,
		LastLoginAt:      &lastLogin,
		LastLoginIp:      &lastLoginIp,
		LastLoginCountry: &lastLoginCountry,
		LastLoginCity:    &lastLoginCity,
		CreatedAt:        createdAt,
		UpdatedAt:        &updatedAt,
	}
}

package entities

import (
	"time"

	"github.com/google/uuid"
)

type account struct {
	id               string
	name             string
	email            string
	pass             string
	phone            string
	is2faEnabled     bool
	isEmailVerified  bool
	lastLogin        *time.Time
	lastLoginIp      *string
	lastLoginCountry *string
	lastLoginCity    *string
	createdAt        time.Time
	updatedAt        *time.Time
}

type IAccount interface {
	GetId() string
	GetName() string
	GetEmail() string
	GetPass() string
	GetPhone() string
	Get2faEnabled() bool
	GetIsEmailVerified() bool
	GetLastLogin() time.Time
	GetLastLoginIp() string
	GetLastLoginCountry() string
	GetLastLoginCity() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

func NewAccount(name, email, pass, phone string) IAccount {
	return &account{
		id:               uuid.New().String(),
		name:             name,
		email:            email,
		pass:             pass,
		phone:            phone,
		is2faEnabled:     false,
		isEmailVerified:  false,
		lastLogin:        nil,
		lastLoginIp:      nil,
		lastLoginCountry: nil,
		lastLoginCity:    nil,
		createdAt:        time.Now(),
		updatedAt:        nil,
	}
}

func NewExistingAccount(id, name, email, pass, phone string, is2faEnabled, isEmailVerified bool, lastLogin time.Time, lastLoginIp, lastLoginCountry, lastLoginCity string, createdAt time.Time, updatedAt time.Time) IAccount {
	return &account{
		id:               id,
		name:             name,
		email:            email,
		pass:             pass,
		phone:            phone,
		is2faEnabled:     is2faEnabled,
		isEmailVerified:  isEmailVerified,
		lastLogin:        &lastLogin,
		lastLoginIp:      &lastLoginIp,
		lastLoginCountry: &lastLoginCountry,
		lastLoginCity:    &lastLoginCity,
		createdAt:        createdAt,
		updatedAt:        &updatedAt,
	}
}

func (u *account) GetId() string {
	return u.id
}

func (u *account) GetName() string {
	return u.name
}

func (u *account) GetEmail() string {
	return u.email
}

func (u *account) GetPass() string {
	return u.pass
}

func (u *account) GetPhone() string {
	return u.phone
}

func (u *account) Get2faEnabled() bool {
	return u.is2faEnabled
}

func (u *account) GetIsEmailVerified() bool {
	return u.isEmailVerified
}

func (u *account) GetLastLogin() time.Time {
	return *u.lastLogin
}

func (u *account) GetLastLoginIp() string {
	return *u.lastLoginIp
}

func (u *account) GetLastLoginCountry() string {
	return *u.lastLoginCountry
}

func (u *account) GetLastLoginCity() string {
	return *u.lastLoginCity
}

func (u *account) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *account) GetUpdatedAt() time.Time {
	return *u.updatedAt
}

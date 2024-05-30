package entities

import (
	"time"

	"github.com/google/uuid"
)

type user struct {
	id               string
	name             string
	email            string
	pass             string
	phone            string
	is2faEnabled     bool
	lastLogin        *time.Time
	lastLoginIp      *string
	lastLoginCountry *string
	lastLoginCity    *string
	createdAt        time.Time
	updatedAt        *time.Time
}

type IUser interface {
	GetId() string
	GetName() string
	GetEmail() string
	GetPass() string
	GetPhone() string
	Get2faEnabled() bool
	GetLastLogin() time.Time
	GetLastLoginIp() string
	GetLastLoginCountry() string
	GetLastLoginCity() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

func NewUser(name, email, pass, phone string) IUser {
	return &user{
		id:               uuid.New().String(),
		name:             name,
		email:            email,
		pass:             pass,
		phone:            phone,
		is2faEnabled:     false,
		lastLogin:        nil,
		lastLoginIp:      nil,
		lastLoginCountry: nil,
		lastLoginCity:    nil,
		createdAt:        time.Now(),
		updatedAt:        nil,
	}
}

func NewExistingUser(id, name, email, pass, phone string, is2faEnabled bool, lastLogin time.Time, lastLoginIp, lastLoginCountry, lastLoginCity string, createdAt time.Time, updatedAt time.Time) IUser {
	return &user{
		id:               id,
		name:             name,
		email:            email,
		pass:             pass,
		phone:            phone,
		is2faEnabled:     is2faEnabled,
		lastLogin:        &lastLogin,
		lastLoginIp:      &lastLoginIp,
		lastLoginCountry: &lastLoginCountry,
		lastLoginCity:    &lastLoginCity,
		createdAt:        createdAt,
		updatedAt:        &updatedAt,
	}
}

func (u *user) GetId() string {
	return u.id
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) GetEmail() string {
	return u.email
}

func (u *user) GetPass() string {
	return u.pass
}

func (u *user) GetPhone() string {
	return u.phone
}

func (u *user) Get2faEnabled() bool {
	return u.is2faEnabled
}

func (u *user) GetLastLogin() time.Time {
	return *u.lastLogin
}

func (u *user) GetLastLoginIp() string {
	return *u.lastLoginIp
}

func (u *user) GetLastLoginCountry() string {
	return *u.lastLoginCountry
}

func (u *user) GetLastLoginCity() string {
	return *u.lastLoginCity
}

func (u *user) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *user) GetUpdatedAt() time.Time {
	return *u.updatedAt
}

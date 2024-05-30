package entities

import (
	"time"

	"github.com/google/uuid"
)

type user struct {
	id        string
	name      string
	email     string
	pass      string
	createdAt time.Time
	updatedAt *time.Time
}

type IUser interface {
	GetId() string
	GetName() string
	GetEmail() string
	GetPass() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

func NewUser(name, email, pass string) IUser {
	return &user{
		id:        uuid.New().String(),
		name:      name,
		email:     email,
		pass:      pass,
		createdAt: time.Now(),
		updatedAt: nil,
	}
}

func NewExistingUser(id, name, email, pass string, createdAt time.Time, updatedAt *time.Time) IUser {
	return &user{
		id:        id,
		name:      name,
		email:     email,
		pass:      pass,
		createdAt: createdAt,
		updatedAt: updatedAt,
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

func (u *user) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *user) GetUpdatedAt() time.Time {
	return *u.updatedAt
}

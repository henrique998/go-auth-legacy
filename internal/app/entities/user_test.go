package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_it_should_create_a_new_user(t *testing.T) {
	assert := assert.New(t)

	user := NewUser("Henrique", "henriquemonteiro@gmail.com", "12345")

	assert.Equal("Henrique", user.GetName())
	assert.Equal("henriquemonteiro@gmail.com", user.GetEmail())
	assert.Equal("12345", user.GetPass())
}

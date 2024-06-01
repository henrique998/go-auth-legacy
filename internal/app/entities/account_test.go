package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_it_should_create_a_new_user(t *testing.T) {
	assert := assert.New(t)

	user := NewAccount("Henrique", "henriquemonteiro@gmail.com", "12345", "999999999")

	assert.Equal("Henrique", user.Name)
	assert.Equal("henriquemonteiro@gmail.com", user.Email)
	assert.Equal("12345", user.Pass)
}

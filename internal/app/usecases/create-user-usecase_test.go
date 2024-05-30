package usecases

import (
	"testing"

	"github.com/henrique998/go-setup/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserUsecase_it_should_be_able_to_create_with_valid_data(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUsersRepository(ctrl)
	sut := CreateUserUseCase{
		Repo: repo,
	}

	repo.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	repo.EXPECT().Create(gomock.Any()).Return(nil)

	err := sut.Execute("Henrique", "henriquemonteiro@gmail.com", "12345")

	assert.Nil(err)
}

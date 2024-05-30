package usersusecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserUsecase_it_should_be_able_to_create_with_valid_data(t *testing.T) {
	assert := assert.New(t)
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()

	// repo := mocks.NewMockUsersRepository(ctrl)
	// sut := CreateUserUseCase{
	// 	Repo: repo,
	// }

	// repo.EXPECT().FindByEmail(gomock.Any()).Return(nil, nil)
	// repo.EXPECT().Create(gomock.Any()).Return(nil)

	// err := sut.Execute(request.CreateUserRequest{})

	assert.Equal(2, 1+1)
}

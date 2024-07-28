package accountsusecases

import (
	"net/http"
	"testing"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserUsecase(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountsRepo := mocks.NewMockAccountsRepository(ctrl)
	mockVTRepo := mocks.NewMockVerificationTokensRepository(ctrl)
	mockEmailProvider := mocks.NewMockEmailProvider(ctrl)

	usecase := CreateAccountUseCase{
		Repo:          mockAccountsRepo,
		VTRepo:        mockVTRepo,
		EmailProvider: mockEmailProvider,
	}

	t.Run("should not be able to create an account with an used email", func(t *testing.T) {
		req := request.CreateAccountRequest{
			Name:  "Jhon doe",
			Email: "jhondoe@gmail.com",
			Pass:  "123456",
			Phone: "9999999",
		}

		account := entities.NewAccount(req.Name, req.Email, req.Pass, req.Phone, "")

		mockAccountsRepo.EXPECT().FindByEmail(req.Email).Return(account)

		err := usecase.Execute(req)

		assert.NotNil(err)
		assert.Equal("account already exists", err.GetMessage())
		assert.Equal(http.StatusBadRequest, err.GetStatus())
	})

	t.Run("It should be able to create an account with valid data", func(t *testing.T) {
		req := request.CreateAccountRequest{
			Name:  "Jhon doe",
			Email: "jhondoe@gmail.com",
			Pass:  "123456",
			Phone: "9999999",
		}

		mockAccountsRepo.EXPECT().FindByEmail(req.Email).Return(nil)
		mockAccountsRepo.EXPECT().Create(gomock.Any()).Return(nil)
		mockVTRepo.EXPECT().Create(gomock.Any()).Return(nil)
		mockEmailProvider.EXPECT().SendMail(req.Email, "Account verification.", gomock.Any())

		err := usecase.Execute(req)

		assert.Nil(err)
	})
}

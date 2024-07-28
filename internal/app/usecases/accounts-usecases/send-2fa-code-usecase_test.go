package accountsusecases

import (
	"net/http"
	"os"
	"testing"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSend2faCodeUseCase(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountsRepo := mocks.NewMockAccountsRepository(ctrl)
	mockVTRepo := mocks.NewMockVerificationTokensRepository(ctrl)
	mockTFaProvider := mocks.NewMockTwoFactorAuthProvider(ctrl)

	usecase := Send2faCodeUseCase{
		Repo:                  mockAccountsRepo,
		VTRepo:                mockVTRepo,
		TwoFactorAuthProvider: mockTFaProvider,
	}

	t.Run("It should not be able to send 2fa code if account not exists", func(t *testing.T) {
		accountId := "invalid-account-id"

		mockAccountsRepo.EXPECT().FindById(accountId).Return(nil)

		err := usecase.Execute(accountId)

		assert.NotNil(err)
		assert.Equal("account does not exists", err.GetMessage())
		assert.Equal(http.StatusNotFound, err.GetStatus())
	})

	t.Run("It should not be able to send 2fa code if phone not exists", func(t *testing.T) {
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "123456", "", "")

		mockAccountsRepo.EXPECT().FindById(account.ID).Return(account)

		err := usecase.Execute(account.ID)

		assert.NotNil(err)
		assert.Equal("account must have an phone number to complete 2fa proccess", err.GetMessage())
		assert.Equal(http.StatusUnauthorized, err.GetStatus())
	})

	t.Run("It should be able to send 2fa code", func(t *testing.T) {
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "123456", "999999999", "")

		mockAccountsRepo.EXPECT().FindById(account.ID).Return(account)
		mockVTRepo.EXPECT().Create(gomock.Any()).Return(nil)

		mockTFaProvider.EXPECT().Send(os.Getenv("TWILIO_FROM_PHONE_NUMBER"), *account.Phone, gomock.Any()).Return(nil)

		err := usecase.Execute(account.ID)

		assert.Nil(err)
	})
}

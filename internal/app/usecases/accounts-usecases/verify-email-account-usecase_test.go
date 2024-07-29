package accountsusecases

import (
	"net/http"
	"testing"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/infra/utils"
	"github.com/henrique998/go-auth/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestVerifyEmailAccountUseCase(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountsRepo := mocks.NewMockAccountsRepository(ctrl)
	mockVTRepo := mocks.NewMockVerificationCodesRepository(ctrl)

	sut := VerifyEmailUseCase{
		Repo:   mockAccountsRepo,
		VTRepo: mockVTRepo,
	}

	t.Run("It should not be able to complete email verification flow if code has expired", func(t *testing.T) {
		codeStr, _ := utils.GenerateCode(10)
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "123456", "999999999", "")
		code := entities.NewVerificationCode(codeStr, account.ID, time.Now().Add(-1*time.Hour))

		mockVTRepo.EXPECT().FindByValue(codeStr).Return(code)

		err := sut.Execute(code.Value)

		assert.NotNil(err)
		assert.Equal("verification code has expired", err.GetMessage())
		assert.Equal(http.StatusUnauthorized, err.GetStatus())
	})

	t.Run("It should not be able to complete 2fa flow if already active", func(t *testing.T) {
		accountId := utils.GenerateUUID()
		account := entities.NewExistingAccount(
			accountId,
			"jhon doe",
			"jhondoe@gmail.com",
			"123456",
			"999999999",
			"",
			true,
			true,
			time.Now().Add(-10*time.Minute),
			"0.0.0.0",
			"br",
			"sp",
			time.Now().Add(-10*(time.Hour*24*10)),
			time.Now().Add(-5*time.Hour),
		)
		codeStr, _ := utils.GenerateCode(10)
		code := entities.NewVerificationCode(codeStr, accountId, time.Now().Add(10*time.Hour))

		mockVTRepo.EXPECT().FindByValue(codeStr).Return(code)
		mockAccountsRepo.EXPECT().FindById(accountId).Return(account)

		err := sut.Execute(codeStr)

		assert.NotNil(err)
		assert.Equal("the email has already been verified", err.GetMessage())
		assert.Equal(http.StatusUnauthorized, err.GetStatus())
	})

	t.Run("It should be able to complete email verification flow", func(t *testing.T) {
		codeStr, _ := utils.GenerateCode(10)
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "123456", "999999999", "")
		code := entities.NewVerificationCode(codeStr, account.ID, time.Now().Add(10*time.Minute))

		mockVTRepo.EXPECT().FindByValue(gomock.Any()).Return(code)
		mockVTRepo.EXPECT().Delete(gomock.Any()).Return(nil)
		mockAccountsRepo.EXPECT().FindById(gomock.Any()).Return(account)
		mockAccountsRepo.EXPECT().Update(gomock.Any()).Return(nil)

		err := sut.Execute(code.Value)

		assert.Nil(err)
	})
}

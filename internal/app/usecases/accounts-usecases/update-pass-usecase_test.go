package accountsusecases

import (
	"net/http"
	"testing"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/infra/utils"
	"github.com/henrique998/go-auth/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUpdatePassUseCase(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountsRepo := mocks.NewMockAccountsRepository(ctrl)
	mockVTRepo := mocks.NewMockVerificationCodesRepository(ctrl)

	usecase := UpdatePassUsecase{
		Repo:   mockAccountsRepo,
		VTRepo: mockVTRepo,
	}

	t.Run("It should not be able to update pass if code not exists", func(t *testing.T) {
		req := request.NewPassRequest{
			Code:                "",
			NewPass:             "updated-pass",
			NewPassConfirmation: "updated-pass",
		}

		mockVTRepo.EXPECT().FindByValue(req.Code).Return(nil)

		err := usecase.Execute(req)

		assert.NotNil(err)
		assert.Equal("code not found", err.GetMessage())
		assert.Equal(http.StatusNotFound, err.GetStatus())
	})

	t.Run("It should not be able to update pass if code has expired", func(t *testing.T) {
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "123456", "999999999", "")
		codeStr, _ := utils.GenerateCode(10)

		code := entities.NewVerificationCode(codeStr, account.ID, time.Now().Add(-1*time.Hour))

		req := request.NewPassRequest{
			Code:                codeStr,
			NewPass:             "updated-pass",
			NewPassConfirmation: "updated-pass",
		}

		mockAccountsRepo.EXPECT().FindById(account.ID).Return(account)
		mockVTRepo.EXPECT().FindByValue(req.Code).Return(code)

		err := usecase.Execute(req)

		assert.NotNil(err)
		assert.Equal("code has expired", err.GetMessage())
		assert.Equal(http.StatusUnauthorized, err.GetStatus())
	})

	t.Run("It should not be able to update pass if pass and confirmation not equals", func(t *testing.T) {
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "123456", "999999999", "")
		codeStr, _ := utils.GenerateCode(10)

		code := entities.NewVerificationCode(codeStr, account.ID, time.Now().Add(10*time.Hour))

		req := request.NewPassRequest{
			Code:                codeStr,
			NewPass:             "updated-pass",
			NewPassConfirmation: "diferent-updated-pass",
		}

		mockAccountsRepo.EXPECT().FindById(account.ID).Return(account)
		mockVTRepo.EXPECT().FindByValue(req.Code).Return(code)

		err := usecase.Execute(req)

		assert.NotNil(err)
		assert.Equal("new password and confirmation must be equals", err.GetMessage())
		assert.Equal(http.StatusBadRequest, err.GetStatus())
	})

	t.Run("It should not be able to update pass if new pass less than 6 characters", func(t *testing.T) {
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "123456", "999999999", "")
		codeStr, _ := utils.GenerateCode(10)

		code := entities.NewVerificationCode(codeStr, account.ID, time.Now().Add(10*time.Hour))

		req := request.NewPassRequest{
			Code:                codeStr,
			NewPass:             "short",
			NewPassConfirmation: "short",
		}

		mockAccountsRepo.EXPECT().FindById(account.ID).Return(account)
		mockVTRepo.EXPECT().FindByValue(req.Code).Return(code)

		err := usecase.Execute(req)

		assert.NotNil(err)
		assert.Equal("new password must contain 6 or more characters", err.GetMessage())
		assert.Equal(http.StatusBadRequest, err.GetStatus())
	})

	t.Run("It should not be able to update pass if earlier pass and new pass equals", func(t *testing.T) {
		hashedPass, _ := utils.HashPass("123456")

		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", hashedPass, "999999999", "")
		codeStr, _ := utils.GenerateCode(10)

		code := entities.NewVerificationCode(codeStr, account.ID, time.Now().Add(10*time.Hour))

		req := request.NewPassRequest{
			Code:                codeStr,
			NewPass:             "123456",
			NewPassConfirmation: "123456",
		}

		mockAccountsRepo.EXPECT().FindById(account.ID).Return(account)
		mockVTRepo.EXPECT().FindByValue(req.Code).Return(code)

		err := usecase.Execute(req)

		assert.NotNil(err)
		assert.Equal("new password cannot be the same as the previous one", err.GetMessage())
		assert.Equal(http.StatusBadRequest, err.GetStatus())
	})

	t.Run("It should be able to update pass", func(t *testing.T) {
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "123456", "999999999", "")
		codeStr, _ := utils.GenerateCode(10)
		now := time.Now()

		code := entities.NewVerificationCode(codeStr, account.ID, now.Add(10*time.Hour))

		req := request.NewPassRequest{
			Code:                codeStr,
			NewPass:             "updated-pass",
			NewPassConfirmation: "updated-pass",
		}

		mockAccountsRepo.EXPECT().FindById(account.ID).Return(account)
		mockAccountsRepo.EXPECT().Update(gomock.Any()).Return(nil)
		mockVTRepo.EXPECT().FindByValue(req.Code).Return(code)
		mockVTRepo.EXPECT().Delete(code.ID).Return(nil)

		err := usecase.Execute(req)

		assert.Nil(err)
	})
}

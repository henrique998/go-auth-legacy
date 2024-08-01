package sessionusecases

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/infra/utils"
	"github.com/henrique998/go-auth/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLoginWithMagicLinkUseCase(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountsRepo := mocks.NewMockAccountsRepository(ctrl)
	mockMLRepo := mocks.NewMockMagicLinksRepository(ctrl)
	mockDevicesRepo := mocks.NewMockDevicesRepository(ctrl)
	mockATProvider := mocks.NewMockAuthTokensProvider(ctrl)
	mockEmailProvider := mocks.NewMockEmailProvider(ctrl)

	sut := LoginWithMagicLinkUseCase{
		Repo:          mockAccountsRepo,
		MLRepo:        mockMLRepo,
		DevicesRepo:   mockDevicesRepo,
		ATProvider:    mockATProvider,
		EmailProvider: mockEmailProvider,
	}

	t.Run("It should not be able to login if magic link does not exists", func(t *testing.T) {
		req := request.LoginWithMagicLinkRequest{
			Code:      "invalid-code",
			IP:        "fake-ip",
			UserAgent: "fake-user-agent",
		}

		mockMLRepo.EXPECT().FindByValue(req.Code).Return(nil)

		accessToken, refreshToken, err := sut.Execute(req)

		assert.Empty(accessToken)
		assert.Empty(refreshToken)
		assert.NotNil(err)
		assert.Equal("Code not found", err.GetMessage())
		assert.Equal(http.StatusNotFound, err.GetStatus())
	})

	t.Run("It should not be able to login if magic link has expired", func(t *testing.T) {
		code, _ := utils.GenerateCode(32)

		req := request.LoginWithMagicLinkRequest{
			Code:      code,
			IP:        "fake-ip",
			UserAgent: "fake-user-agent",
		}

		expiresAt := time.Now().Add(-15 * time.Minute)
		ml := entities.NewMagicLink("account-id", code, expiresAt)

		mockMLRepo.EXPECT().FindByValue(req.Code).Return(ml)

		accessToken, refreshToken, err := sut.Execute(req)

		assert.Empty(accessToken)
		assert.Empty(refreshToken)
		assert.NotNil(err)
		assert.Equal("Code has expired", err.GetMessage())
		assert.Equal(http.StatusUnauthorized, err.GetStatus())
	})

	t.Run("It should not be able to login if account does not exists", func(t *testing.T) {
		code, _ := utils.GenerateCode(32)

		req := request.LoginWithMagicLinkRequest{
			Code:      code,
			IP:        "fake-ip",
			UserAgent: "fake-user-agent",
		}

		expiresAt := time.Now().Add(15 * time.Minute)
		ml := entities.NewMagicLink("account-id", code, expiresAt)

		mockMLRepo.EXPECT().FindByValue(req.Code).Return(ml)
		mockAccountsRepo.EXPECT().FindById(gomock.Any()).Return(nil)

		accessToken, refreshToken, err := sut.Execute(req)

		assert.Empty(accessToken)
		assert.Empty(refreshToken)
		assert.NotNil(err)
		assert.Equal("Account not found", err.GetMessage())
		assert.Equal(http.StatusNotFound, err.GetStatus())
	})

	t.Run("It should not be able to login if account does not exists", func(t *testing.T) {
		code, _ := utils.GenerateCode(32)

		req := request.LoginWithMagicLinkRequest{
			Code:      code,
			IP:        "fake-ip",
			UserAgent: "fake-user-agent",
		}

		expiresAt := time.Now().Add(15 * time.Minute)
		ml := entities.NewMagicLink("account-id", code, expiresAt)

		mockMLRepo.EXPECT().FindByValue(req.Code).Return(ml)
		mockAccountsRepo.EXPECT().FindById(gomock.Any()).Return(nil)

		accessToken, refreshToken, err := sut.Execute(req)

		assert.Empty(accessToken)
		assert.Empty(refreshToken)
		assert.NotNil(err)
		assert.Equal("Account not found", err.GetMessage())
		assert.Equal(http.StatusNotFound, err.GetStatus())
	})

	t.Run("It should not be able to login with magic link", func(t *testing.T) {
		code, _ := utils.GenerateCode(32)

		req := request.LoginWithMagicLinkRequest{
			Code:      code,
			IP:        "fake-ip",
			UserAgent: "fake-user-agent",
		}

		expiresAt := time.Now().Add(15 * time.Minute)
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "", "", "")
		ml := entities.NewMagicLink(account.ID, code, expiresAt)
		accessTokenStr, _ := utils.GenerateJWTToken(account.ID, expiresAt, os.Getenv("JWT_SECRET"))
		refreshTokenStr, _ := utils.GenerateJWTToken(account.ID, time.Now().Add(1*(time.Hour*24*30)), os.Getenv("JWT_SECRET"))

		mockMLRepo.EXPECT().FindByValue(req.Code).Return(ml)
		mockAccountsRepo.EXPECT().FindById(account.ID).Return(account)
		mockATProvider.EXPECT().GenerateAuthTokens(account.ID).Return(accessTokenStr, refreshTokenStr, nil)
		mockDevicesRepo.EXPECT().FindByIpAndAccountId(gomock.Any(), gomock.Any()).Return(nil)
		mockDevicesRepo.EXPECT().Create(gomock.Any()).Return(nil)
		mockAccountsRepo.EXPECT().Update(gomock.Any()).Return(nil)
		mockMLRepo.EXPECT().Delete(ml.ID).Return(nil)

		accessToken, refreshToken, err := sut.Execute(req)

		assert.Equal(accessTokenStr, accessToken)
		assert.Equal(refreshTokenStr, refreshToken)
		assert.Nil(err)
	})
}

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

func TestLoginWithCredentialsUseCase(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountsRepo := mocks.NewMockAccountsRepository(ctrl)
	mockDevicesRepo := mocks.NewMockDevicesRepository(ctrl)
	mockLARepo := mocks.NewMockLoginAttemptsRepository(ctrl)
	mockEmailProvider := mocks.NewMockEmailProvider(ctrl)
	mockATProvider := mocks.NewMockAuthTokensProvider(ctrl)
	mockGLProvider := mocks.NewMockGeoLocationProvider(ctrl)

	usecase := LoginWithCredentialsUseCase{
		Repo:          mockAccountsRepo,
		DevicesRepo:   mockDevicesRepo,
		LARepository:  mockLARepo,
		EmailProvider: mockEmailProvider,
		AtProvider:    mockATProvider,
		GLProvider:    mockGLProvider,
	}

	t.Run("It should not be able to login with credentials if with wrong email", func(t *testing.T) {
		req := request.LoginWithCredentialsRequest{
			Email:     "invalid-email",
			Pass:      "123456",
			IP:        "0.0.0.0",
			UserAgent: "user-agent",
		}

		mockAccountsRepo.EXPECT().FindByEmail(req.Email).Return(nil)
		mockLARepo.EXPECT().Create(gomock.Any()).Return(nil)

		accessToken, refreshToken, err := usecase.Execute(req)

		assert.Empty(accessToken)
		assert.Empty(refreshToken)
		assert.NotNil(err)
		assert.Equal("email or password incorrect!", err.GetMessage())
		assert.Equal(http.StatusBadRequest, err.GetStatus())
	})

	t.Run("It should not be able to login with credentials if pass does not exists", func(t *testing.T) {
		req := request.LoginWithCredentialsRequest{
			Email:     "invalid-email",
			Pass:      "123456",
			IP:        "0.0.0.0",
			UserAgent: "user-agent",
		}

		account := entities.NewAccount("jhon doe", req.Email, "", "", "")

		mockAccountsRepo.EXPECT().FindByEmail(req.Email).Return(account)
		mockLARepo.EXPECT().Create(gomock.Any()).Return(nil)

		accessToken, refreshToken, err := usecase.Execute(req)

		assert.Empty(accessToken)
		assert.Empty(refreshToken)
		assert.NotNil(err)
		assert.Equal("Login method not allowed!", err.GetMessage())
		assert.Equal(http.StatusUnauthorized, err.GetStatus())
	})

	t.Run("It should not be able to login with credentials with wrong password", func(t *testing.T) {
		email := "jhondoe@gmail.com"
		hashedPass, _ := utils.HashPass("123456")

		req := request.LoginWithCredentialsRequest{
			Email:     email,
			Pass:      "wrong-pass",
			IP:        "0.0.0.0",
			UserAgent: "user-agent",
		}

		account := entities.NewAccount("jhon doe", email, hashedPass, "", "")

		mockAccountsRepo.EXPECT().FindByEmail(req.Email).Return(account)
		mockLARepo.EXPECT().Create(gomock.Any()).Return(nil)

		accessToken, refreshToken, err := usecase.Execute(req)

		assert.Empty(accessToken)
		assert.Empty(refreshToken)
		assert.NotNil(err)
		assert.Equal("email or password incorrect!", err.GetMessage())
		assert.Equal(http.StatusBadRequest, err.GetStatus())
	})

	t.Run("It should not be able to login with credentials with unveried email", func(t *testing.T) {
		email := "jhondoe@gmail.com"
		pass := "123456"
		hashedPass, _ := utils.HashPass(pass)

		req := request.LoginWithCredentialsRequest{
			Email:     email,
			Pass:      pass,
			IP:        "0.0.0.0",
			UserAgent: "user-agent",
		}

		account := entities.NewExistingAccount("id", "jhon doe", email, hashedPass, "", "", false, false, time.Now().Add(-1*time.Hour), req.IP, "br", "sp", time.Now().Add(-1*time.Hour*24*30), time.Now().Add(-2*time.Hour*24*2))

		mockAccountsRepo.EXPECT().FindByEmail(req.Email).Return(account)
		mockLARepo.EXPECT().Create(gomock.Any()).Return(nil)

		accessToken, refreshToken, err := usecase.Execute(req)

		assert.Empty(accessToken)
		assert.Empty(refreshToken)
		assert.NotNil(err)
		assert.Equal("Only verified accounts can log in!", err.GetMessage())
		assert.Equal(http.StatusUnauthorized, err.GetStatus())
	})

	t.Run("It should be able to login with credentials sending notification email because diferent geolocation info", func(t *testing.T) {
		email := "jhondoe@gmail.com"
		pass := "123456"
		hashedPass, _ := utils.HashPass(pass)

		req := request.LoginWithCredentialsRequest{
			Email:     email,
			Pass:      pass,
			IP:        "0.0.0.0",
			UserAgent: "user-agent",
		}

		account := entities.NewExistingAccount("id", "jhon doe", email, hashedPass, "", "", false, true, time.Now().Add(-1*time.Hour), req.IP, "br", "sp", time.Now().Add(-1*time.Hour*24*30), time.Now().Add(-2*time.Hour*24*2))
		accessTokenStr, _ := utils.GenerateJWTToken(account.ID, time.Now().Add(15*time.Minute), os.Getenv("JWT_SECRET"))
		refreshTokenStr, _ := utils.GenerateJWTToken(account.ID, time.Now().Add(1*time.Hour*24*30), os.Getenv("JWT_SECRET"))

		mockAccountsRepo.EXPECT().FindByEmail(req.Email).Return(account)
		mockATProvider.EXPECT().GenerateAuthTokens(account.ID).Return(accessTokenStr, refreshTokenStr, nil)
		mockGLProvider.EXPECT().GetInfo(req.IP).Return("br", "rj", nil)
		mockEmailProvider.EXPECT().SendMail(req.Email, "login suspeito", gomock.Any()).Return(nil)
		mockDevicesRepo.EXPECT().FindByIpAndAccountId(gomock.Any(), gomock.Any()).Return(nil)
		mockDevicesRepo.EXPECT().Create(gomock.Any()).Return(nil)
		mockAccountsRepo.EXPECT().Update(gomock.Any()).Return(nil)
		mockLARepo.EXPECT().Create(gomock.Any()).Return(nil)

		accessToken, refreshToken, err := usecase.Execute(req)

		assert.Nil(err)
		assert.Equal(accessTokenStr, accessToken)
		assert.Equal(refreshTokenStr, refreshToken)
	})

	t.Run("It should be able to login with credentials", func(t *testing.T) {
		email := "jhondoe@gmail.com"
		pass := "123456"
		hashedPass, _ := utils.HashPass(pass)

		req := request.LoginWithCredentialsRequest{
			Email:     email,
			Pass:      pass,
			IP:        "0.0.0.0",
			UserAgent: "user-agent",
		}

		account := entities.NewExistingAccount("id", "jhon doe", email, hashedPass, "", "", false, true, time.Now().Add(-1*time.Hour), req.IP, "br", "sp", time.Now().Add(-1*time.Hour*24*30), time.Now().Add(-2*time.Hour*24*2))
		accessTokenStr, _ := utils.GenerateJWTToken(account.ID, time.Now().Add(15*time.Minute), os.Getenv("JWT_SECRET"))
		refreshTokenStr, _ := utils.GenerateJWTToken(account.ID, time.Now().Add(1*time.Hour*24*30), os.Getenv("JWT_SECRET"))

		mockAccountsRepo.EXPECT().FindByEmail(req.Email).Return(account)
		mockATProvider.EXPECT().GenerateAuthTokens(account.ID).Return(accessTokenStr, refreshTokenStr, nil)
		mockGLProvider.EXPECT().GetInfo(req.IP).Return("br", "sp", nil)
		mockDevicesRepo.EXPECT().FindByIpAndAccountId(gomock.Any(), gomock.Any()).Return(nil)
		mockDevicesRepo.EXPECT().Create(gomock.Any()).Return(nil)
		mockAccountsRepo.EXPECT().Update(gomock.Any()).Return(nil)
		mockLARepo.EXPECT().Create(gomock.Any()).Return(nil)

		accessToken, refreshToken, err := usecase.Execute(req)

		assert.Nil(err)
		assert.Equal(accessTokenStr, accessToken)
		assert.Equal(refreshTokenStr, refreshToken)
	})
}

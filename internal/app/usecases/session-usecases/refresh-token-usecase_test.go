package sessionusecases

import (
	"net/http"
	"testing"

	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRefreshTokenUseCase(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRefreshTokensRepo := mocks.NewMockRefreshTokensRepository(ctrl)
	mockATRepo := mocks.NewMockAuthTokensProvider(ctrl)

	sut := RefreshTokenUseCase{
		Repo:   mockRefreshTokensRepo,
		ATRepo: mockATRepo,
	}

	t.Run("It should return an error if the token validation fails", func(t *testing.T) {
		token := "invalid-token"
		mockATRepo.EXPECT().ValidateJWTToken(token).Return("", errors.NewAppError("invalid token", http.StatusUnauthorized))

		accessToken, refreshToken, err := sut.Execute(token)

		assert.Empty(accessToken)
		assert.Empty(refreshToken)
		assert.NotNil(err)
		assert.Equal("invalid token", err.GetMessage())
		assert.Equal(http.StatusUnauthorized, err.GetStatus())
	})

	t.Run("It should return new access and refresh tokens if everything goes well", func(t *testing.T) {
		validToken := "valid.token.here"
		accountId := "account-id"

		newAccessToken := "new.access.token"
		newRefreshToken := "new.refresh.token"

		mockATRepo.EXPECT().ValidateJWTToken(validToken).Return(accountId, nil)
		mockATRepo.EXPECT().GenerateAuthTokens(accountId).Return(newAccessToken, newRefreshToken, nil)
		mockRefreshTokensRepo.EXPECT().Delete(validToken).Return(nil)

		accessToken, refreshToken, err := sut.Execute(validToken)

		assert.Equal(newAccessToken, accessToken)
		assert.Equal(newRefreshToken, refreshToken)
		assert.Nil(err)
	})
}

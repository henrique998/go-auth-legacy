package sessionusecases

import (
	"net/http"
	"testing"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRequestMagicLink(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountsRepo := mocks.NewMockAccountsRepository(ctrl)
	mockMLProvider := mocks.NewMockMagicLinksRepository(ctrl)
	mockEmailProvider := mocks.NewMockEmailProvider(ctrl)

	sut := RequestMagicLinkUseCase{
		Repo:          mockAccountsRepo,
		MLRepo:        mockMLProvider,
		EmailProvider: mockEmailProvider,
	}

	t.Run("It should not be able to send a request if account does not exists", func(t *testing.T) {
		email := "wrong-email"

		mockAccountsRepo.EXPECT().FindByEmail(email).Return(nil)

		err := sut.Execute(email)

		assert.NotNil(err)
		assert.Equal("account not found", err.GetMessage())
		assert.Equal(http.StatusNotFound, err.GetStatus())
	})

	t.Run("It should be able to send request for a magic link", func(t *testing.T) {
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "", "", "")

		mockAccountsRepo.EXPECT().FindByEmail(account.Email).Return(account)
		mockMLProvider.EXPECT().Create(gomock.Any()).Return(nil)
		mockEmailProvider.EXPECT().SendMail(account.Email, "Auth with magic link!", gomock.Any()).Return(nil)

		err := sut.Execute(account.Email)

		assert.Nil(err)
	})
}

package sessionusecases

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/utils"
)

type RequestMagicLinkUseCase struct {
	Repo          contracts.AccountsRepository
	MLRepo        contracts.MagicLinksRepository
	EmailProvider contracts.EmailProvider
}

func (uc *RequestMagicLinkUseCase) Execute(email string) errors.IAppError {
	logger.Info("Init RequestMagicLink UseCase")

	account := uc.Repo.FindByEmail(email)
	if account == nil {
		return errors.NewAppError("email or password incorrect!", http.StatusBadRequest)
	}

	code, err := utils.GenerateToken(32)
	if err != nil {
		logger.Error("Error trying to generate magic link code", err)
		return errors.NewAppError("internal server error!", http.StatusInternalServerError)
	}
	expiresAt := time.Now().Add(15 * time.Minute)

	magicLink := entities.NewMagicLink(account.ID, code, expiresAt)

	err = uc.MLRepo.Create(*magicLink)
	if err != nil {
		logger.Error("Error trying to create magic link", err)
		return errors.NewAppError("internal server error!", http.StatusInternalServerError)
	}

	magicLinkURL := fmt.Sprintf("%s/session/login/magic-link?code=%s", os.Getenv("BASE_URL"), code)
	body := fmt.Sprintf("Click the link to login: %s\nThis link will expire in 15 minutes.", magicLinkURL)

	err = uc.EmailProvider.SendMail(account.Email, "Auth with magic link!", body)
	if err != nil {
		logger.Error("Error trying to send email with magic link", err)
		return errors.NewAppError("internal server error!", http.StatusInternalServerError)
	}

	return nil
}

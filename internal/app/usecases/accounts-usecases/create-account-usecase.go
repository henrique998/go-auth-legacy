package accountsusecases

import (
	"errors"

	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/entities"
	appError "github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type CreateAccountUseCase struct {
	Repo contracts.AccountsRepository
}

func (uc *CreateAccountUseCase) Execute(req request.CreateAccountRequest) appError.IAppError {
	logger.Info("Init CreateUser UseCase")

	user := uc.Repo.FindByEmail(req.Email)

	if user != nil {
		logger.Error("Error while validate user.", errors.New("user already exists"))
		return appError.NewAppError("user already exists", 400)
	}

	data := entities.NewAccount(req.Name, req.Email, req.Pass, req.Phone)

	err := uc.Repo.Create(data)
	if err != nil {
		logger.Error("Error trying to create user.", errors.New(err.GetMessage()))
		return err
	}

	return nil
}

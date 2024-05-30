package usecases

import (
	"errors"

	"github.com/henrique998/go-setup/internal/app/contracts"
	"github.com/henrique998/go-setup/internal/app/entities"
	appError "github.com/henrique998/go-setup/internal/app/errors"
	"github.com/henrique998/go-setup/internal/configs/logger"
)

type CreateUserUseCase struct {
	Repo contracts.UsersRepository
}

func (uc *CreateUserUseCase) Execute(name, email, pass string) appError.IAppError {
	logger.Info("Init CreateUser UseCase")

	user, err := uc.Repo.FindByEmail(email)
	if err != nil {
		logger.Error("Error trying to find user.", errors.New(err.GetMessage()))
		return err
	}

	if user != nil {
		logger.Error("Error while validate user.", errors.New("user already exists"))
		return appError.NewAppError("user already exists", 400)
	}

	data := entities.NewUser(name, email, pass)

	err = uc.Repo.Create(data)
	if err != nil {
		logger.Error("Error trying to create user.", errors.New(err.GetMessage()))
		return err
	}

	return nil
}

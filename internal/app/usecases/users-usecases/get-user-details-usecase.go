package usersusecases

import (
	"errors"

	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/henrique998/go-auth/internal/app/entities"
	appError "github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type GetUserDetailsUseCase struct {
	Repo contracts.UsersRepository
}

func (uc *GetUserDetailsUseCase) Execute(email string) (entities.IUser, appError.IAppError) {
	logger.Info("Init GetUserDetails UseCase")

	user := uc.Repo.FindByEmail(email)

	if user == nil {
		logger.Error("Error trying to find user.", errors.New("user not found!"))
		return nil, appError.NewAppError("user not found!", 404)
	}

	return *user, nil
}

package repositories

import (
	"database/sql"
	"fmt"

	"github.com/henrique998/go-setup/internal/app/entities"
	"github.com/henrique998/go-setup/internal/app/errors"
)

type PGUsersRepository struct {
	Db *sql.DB
}

func (r *PGUsersRepository) Create(user entities.IUser) errors.IAppError {
	fmt.Println(user)

	return nil
}

func (r *PGUsersRepository) FindByEmail(email string) (*entities.IUser, errors.IAppError) {
	u := entities.NewUser("j", "h", "hs")

	return &u, nil
}

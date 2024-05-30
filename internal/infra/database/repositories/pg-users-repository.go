package repositories

import (
	"database/sql"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
	"github.com/henrique998/go-auth/internal/infra/database/mappers"
)

type PGUsersRepository struct {
	Db *sql.DB
}

func (r *PGUsersRepository) Create(user entities.IUser) errors.IAppError {
	sql :=
		`INSERT INTO users (id, name, email, password_hash, phone_number, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.Db.Exec(sql,
		user.GetId(),
		user.GetName(),
		user.GetEmail(),
		user.GetPass(),
		user.GetPhone(),
		user.GetCreatedAt(),
	)
	if err != nil {
		return errors.NewAppError(err.Error(), 400)
	}

	return nil
}

func (r *PGUsersRepository) FindByEmail(email string) *entities.IUser {
	var userData mappers.DbUserData

	query := "SELECT id, name, email, password_hash, phone_number, is_2fa_enabled, last_login, last_login_ip, last_login_country, last_login_city, created_at, updated_at FROM users WHERE email = $1"
	row := r.Db.QueryRow(query, email)

	err := row.Scan(
		&userData.ID,
		&userData.Name,
		&userData.Email,
		&userData.Pass,
		&userData.Phone,
		&userData.Is2faEnabled,
		&userData.LastLogin,
		&userData.LastLoginIp,
		&userData.LastLoginCountry,
		&userData.LastLoginCity,
		&userData.CreatedAt,
		&userData.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		logger.Error("Error trying to find user!", err)
	}

	user := mappers.MapToUser(userData)

	return &user
}

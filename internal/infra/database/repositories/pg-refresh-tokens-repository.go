package repositories

import (
	"database/sql"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type PGRefreshTokensRepository struct {
	Db *sql.DB
}

func (r *PGRefreshTokensRepository) FindByValue(val string) (*entities.RefreshToken, errors.IAppError) {
	var refreshToken entities.RefreshToken

	query := "SELECT id, refresh_token, account_id, expires_at, created_at FROM refresh_tokens WHERE refresh_token = $1"
	row := r.Db.QueryRow(query, val)

	err := row.Scan(&refreshToken.ID, &refreshToken.Value, &refreshToken.AccountId, &refreshToken.ExpiresAt, &refreshToken.CreatedAt)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error trying to find verification token!", err)
		return nil, errors.NewAppError(err.Error(), 500)
	}

	return &refreshToken, nil
}

func (r *PGRefreshTokensRepository) Create(rt entities.RefreshToken) errors.IAppError {
	query := "INSERT INTO refresh_tokens (id, refresh_token, account_id, expires_at, created_at) VALUES($1, $2, $3, $4, $5)"

	_, err := r.Db.Exec(query,
		rt.ID,
		rt.Value,
		rt.AccountId,
		rt.ExpiresAt,
		rt.CreatedAt,
	)
	if err != nil {
		return errors.NewAppError(err.Error(), 400)
	}

	return nil
}

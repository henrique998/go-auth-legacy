package repositories

import (
	"database/sql"

	"github.com/henrique998/go-auth/internal/app/entities"
)

type PGRefreshTokensRepository struct {
	Db *sql.DB
}

func (r *PGRefreshTokensRepository) FindByValue(val string) (*entities.RefreshToken, error) {
	var refreshToken entities.RefreshToken

	query := "SELECT id, refresh_token, account_id, expires_at, created_at FROM refresh_tokens WHERE refresh_token = $1"
	row := r.Db.QueryRow(query, val)

	err := row.Scan(&refreshToken.ID, &refreshToken.Value, &refreshToken.AccountId, &refreshToken.ExpiresAt, &refreshToken.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &refreshToken, nil
}

func (r *PGRefreshTokensRepository) Create(rt entities.RefreshToken) error {
	query := "INSERT INTO refresh_tokens (id, refresh_token, account_id, expires_at, created_at) VALUES($1, $2, $3, $4, $5)"

	_, err := r.Db.Exec(query,
		rt.ID,
		rt.Value,
		rt.AccountId,
		rt.ExpiresAt,
		rt.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PGRefreshTokensRepository) Delete(val string) error {
	query := "DELETE FROM refresh_tokens WHERE refresh_token = $1"

	_, err := r.Db.Exec(query, val)
	if err != nil {
		return err
	}

	return nil
}

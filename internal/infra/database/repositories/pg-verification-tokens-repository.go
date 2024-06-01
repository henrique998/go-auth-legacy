package repositories

import (
	"database/sql"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type PGVerificationTokensRepository struct {
	Db *sql.DB
}

func (r *PGVerificationTokensRepository) FindByValue(val string) (*entities.VerificationToken, errors.IAppError) {
	var verificationToken entities.VerificationToken

	query := "SELECT id, account_id, token, created_at, expires_at FROM verification_codes WHERE token = $1"
	row := r.Db.QueryRow(query, val)

	err := row.Scan(&verificationToken.ID, &verificationToken.AccountId, &verificationToken.Value, &verificationToken.CreatedAt, &verificationToken.ExpiresAt)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("Error trying to find verification token!", err)
		return nil, errors.NewAppError(err.Error(), 500)
	}

	return &verificationToken, nil
}

func (r *PGVerificationTokensRepository) Create(verificationToken entities.VerificationToken) errors.IAppError {
	query := "INSERT INTO verification_codes (id, account_id, token, created_at, expires_at) VALUES($1, $2, $3, $4, $5)"

	_, err := r.Db.Exec(query,
		verificationToken.ID,
		verificationToken.AccountId,
		verificationToken.Value,
		verificationToken.CreatedAt,
		verificationToken.ExpiresAt,
	)
	if err != nil {
		return errors.NewAppError(err.Error(), 400)
	}

	return nil
}

func (r *PGVerificationTokensRepository) Delete(tokenId string) errors.IAppError {
	query := "DELETE FROM verification_codes WHERE id = $1"

	_, err := r.Db.Exec(query, tokenId)
	if err != nil {
		return errors.NewAppError(err.Error(), 400)
	}

	return nil
}

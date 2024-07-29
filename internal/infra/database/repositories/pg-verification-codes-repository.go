package repositories

import (
	"database/sql"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type PGVerificationCodesRepository struct {
	Db *sql.DB
}

func (r *PGVerificationCodesRepository) FindByValue(val string) *entities.VerificationCode {
	var verificationCode entities.VerificationCode

	query := "SELECT id, account_id, token, created_at, expires_at FROM verification_codes WHERE token = $1"
	row := r.Db.QueryRow(query, val)

	err := row.Scan(&verificationCode.ID, &verificationCode.AccountId, &verificationCode.Value, &verificationCode.CreatedAt, &verificationCode.ExpiresAt)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error("Error trying to find verification code", err)
		}
		return nil
	}

	return &verificationCode
}

func (r *PGVerificationCodesRepository) Create(verificationCode entities.VerificationCode) error {
	query := "INSERT INTO verification_codes (id, account_id, token, created_at, expires_at) VALUES($1, $2, $3, $4, $5)"

	_, err := r.Db.Exec(query,
		verificationCode.ID,
		verificationCode.AccountId,
		verificationCode.Value,
		verificationCode.CreatedAt,
		verificationCode.ExpiresAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PGVerificationCodesRepository) Delete(tokenId string) error {
	query := "DELETE FROM verification_codes WHERE id = $1"

	_, err := r.Db.Exec(query, tokenId)
	if err != nil {
		return err
	}

	return nil
}

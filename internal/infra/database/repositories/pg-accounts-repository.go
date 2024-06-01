package repositories

import (
	"database/sql"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type PGAccountsRepository struct {
	Db *sql.DB
}

func (r *PGAccountsRepository) FindById(accountId string) (*entities.Account, errors.IAppError) {
	var account entities.Account

	query := "SELECT id, name, email, password_hash, phone_number, is_2fa_enabled, last_login_at, last_login_ip, last_login_country, last_login_city, created_at, updated_at FROM accounts WHERE id = $1"
	row := r.Db.QueryRow(query, accountId)

	err := row.Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.Pass,
		&account.Phone,
		&account.Is2faEnabled,
		&account.LastLoginAt,
		&account.LastLoginIp,
		&account.LastLoginCountry,
		&account.LastLoginCity,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("Error trying to find account!", err)
		return nil, errors.NewAppError(err.Error(), 500)
	}

	return &account, nil
}

func (r *PGAccountsRepository) FindByEmail(email string) (*entities.Account, errors.IAppError) {
	var account entities.Account

	query := "SELECT id, name, email, password_hash, phone_number, is_2fa_enabled, last_login_at, last_login_ip, last_login_country, last_login_city, created_at, updated_at FROM accounts WHERE email = $1"
	row := r.Db.QueryRow(query, email)

	err := row.Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.Pass,
		&account.Phone,
		&account.Is2faEnabled,
		&account.LastLoginAt,
		&account.LastLoginIp,
		&account.LastLoginCountry,
		&account.LastLoginCity,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("Error trying to find account!", err)
		return nil, errors.NewAppError(err.Error(), 500)
	}

	return &account, nil
}

func (r *PGAccountsRepository) Create(account entities.Account) errors.IAppError {
	query :=
		`INSERT INTO accounts (id, name, email, password_hash, phone_number, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.Db.Exec(query,
		account.ID,
		account.Name,
		account.Email,
		account.Pass,
		account.Phone,
		account.CreatedAt,
	)
	if err != nil {
		return errors.NewAppError(err.Error(), 400)
	}

	return nil
}

func (r *PGAccountsRepository) Update(account *entities.Account) errors.IAppError {
	query := "UPDATE accounts SET is_email_verified = $1 WHERE id = $2"

	_, err := r.Db.Exec(query, account.IsEmailVerified, account.ID)
	if err != nil {
		return errors.NewAppError(err.Error(), 400)
	}

	return nil
}

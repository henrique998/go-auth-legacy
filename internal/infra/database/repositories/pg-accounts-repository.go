package repositories

import (
	"database/sql"

	"github.com/henrique998/go-auth/internal/app/entities"
)

type PGAccountsRepository struct {
	Db *sql.DB
}

func (r *PGAccountsRepository) FindById(accountId string) (*entities.Account, error) {
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
		return nil, err
	}

	return &account, nil
}

func (r *PGAccountsRepository) FindByEmail(email string) (*entities.Account, error) {
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
		return nil, err
	}

	return &account, nil
}

func (r *PGAccountsRepository) Create(account entities.Account) error {
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
		return err
	}

	return nil
}

func (r *PGAccountsRepository) Update(account entities.Account) error {
	query := "UPDATE accounts SET name = $1, email = $2, password_hash = $3, phone_number = $4, is_2fa_enabled = $5, is_email_verified = $6, last_login_at = $7, last_login_ip = $8, last_login_country = $9, last_login_city = $10, updated_at = $11 WHERE id = $12"

	_, err := r.Db.Exec(
		query,
		account.Name,
		account.Email,
		account.Pass,
		account.Phone,
		account.Is2faEnabled,
		account.IsEmailVerified,
		account.LastLoginAt,
		account.LastLoginIp,
		account.LastLoginCountry,
		account.LastLoginCity,
		account.UpdatedAt,
		account.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

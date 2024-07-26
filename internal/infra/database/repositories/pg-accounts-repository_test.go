package repositories

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/stretchr/testify/assert"
)

func TestPGAccountsRepository_FindById(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	accountId := "fake-account-id"
	hashedPass := "hashed-pass"
	phone := "999999999"
	providerId := ""

	accountData := entities.Account{
		ID:               accountId,
		Name:             "jhon",
		Email:            "jhondoe@email.com",
		Pass:             &hashedPass,
		Phone:            &phone,
		ProviderId:       &providerId,
		Is2faEnabled:     false,
		IsEmailVerified:  false,
		LastLoginAt:      nil,
		LastLoginIp:      nil,
		LastLoginCountry: nil,
		LastLoginCity:    nil,
		CreatedAt:        time.Now(),
		UpdatedAt:        nil,
	}

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "name", "email", "password_hash", "phone_number", "provider_id", "is_2fa_enabled", "last_login_at", "last_login_ip", "last_login_country", "last_login_city", "created_at", "updated_at",
		}).AddRow(
			accountData.ID,
			accountData.Name,
			accountData.Email,
			accountData.Pass,
			accountData.Phone,
			accountData.ProviderId,
			accountData.Is2faEnabled,
			accountData.LastLoginAt,
			accountData.LastLoginIp,
			accountData.LastLoginCountry,
			accountData.LastLoginCity,
			accountData.CreatedAt,
			accountData.UpdatedAt,
		)

		mock.ExpectQuery("SELECT id, name, email, password_hash, phone_number, provider_id, is_2fa_enabled, last_login_at, last_login_ip, last_login_country, last_login_city, created_at, updated_at FROM accounts WHERE id = \\$1 LIMIT 1").WithArgs(accountId).WillReturnRows(rows)

		repo := PGAccountsRepository{
			Db: db,
		}

		account := repo.FindById(accountId)

		assert.NotNil(t, account)
		assert.Equal(accountId, account.ID)
		assert.Equal(accountData.Email, account.Email)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("ErrorNotNoRows", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, name, email, password_hash, phone_number, provider_id, is_2fa_enabled, last_login_at, last_login_ip, last_login_country, last_login_city, created_at, updated_at FROM accounts WHERE id = \\$1 LIMIT 1").
			WithArgs(accountId).
			WillReturnError(errors.New("some other error"))

		repo := PGAccountsRepository{
			Db: db,
		}

		account := repo.FindById(accountId)

		assert.Nil(account)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

}

func TestPGAccountsRepository_FindByEmail(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	email := "jhondoe@email.com"
	hashedPass := "hashed-pass"
	phone := "999999999"
	providerId := ""

	accountData := entities.Account{
		ID:               "fake-account-id",
		Name:             "jhon",
		Email:            email,
		Pass:             &hashedPass,
		Phone:            &phone,
		ProviderId:       &providerId,
		Is2faEnabled:     false,
		IsEmailVerified:  false,
		LastLoginAt:      nil,
		LastLoginIp:      nil,
		LastLoginCountry: nil,
		LastLoginCity:    nil,
		CreatedAt:        time.Now(),
		UpdatedAt:        nil,
	}

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "name", "email", "password_hash", "phone_number", "provider_id", "is_2fa_enabled", "last_login_at", "last_login_ip", "last_login_country", "last_login_city", "created_at", "updated_at",
		}).AddRow(
			accountData.ID,
			accountData.Name,
			accountData.Email,
			accountData.Pass,
			accountData.Phone,
			accountData.ProviderId,
			accountData.Is2faEnabled,
			accountData.LastLoginAt,
			accountData.LastLoginIp,
			accountData.LastLoginCountry,
			accountData.LastLoginCity,
			accountData.CreatedAt,
			accountData.UpdatedAt,
		)

		mock.ExpectQuery("SELECT id, name, email, password_hash, phone_number, provider_id, is_2fa_enabled, last_login_at, last_login_ip, last_login_country, last_login_city, created_at, updated_at FROM accounts WHERE email = \\$1 LIMIT 1").WithArgs(email).WillReturnRows(rows)

		repo := PGAccountsRepository{
			Db: db,
		}

		account := repo.FindByEmail(email)

		assert.NotNil(t, account)
		assert.Equal(email, account.Email)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("ErrorNotNoRows", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, name, email, password_hash, phone_number, provider_id, is_2fa_enabled, last_login_at, last_login_ip, last_login_country, last_login_city, created_at, updated_at FROM accounts WHERE email = \\$1 LIMIT 1").
			WithArgs(email).
			WillReturnError(errors.New("some other error"))

		repo := PGAccountsRepository{
			Db: db,
		}

		account := repo.FindByEmail(email)

		assert.Nil(account)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGAccountsRepository_Create(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	hashedPass := "hashed-pass"
	phone := "999999999"
	providerId := ""

	repo := PGAccountsRepository{Db: db}

	accountData := entities.Account{
		ID:               "fake-account-id",
		Name:             "jhon",
		Email:            "jhondoe@email.com",
		Pass:             &hashedPass,
		Phone:            &phone,
		ProviderId:       &providerId,
		Is2faEnabled:     false,
		IsEmailVerified:  false,
		LastLoginAt:      nil,
		LastLoginIp:      nil,
		LastLoginCountry: nil,
		LastLoginCity:    nil,
		CreatedAt:        time.Now(),
		UpdatedAt:        nil,
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO accounts \(id, name, email, password_hash, phone_number, provider_id, created_at\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7\)`).
			WithArgs(
				accountData.ID,
				accountData.Name,
				accountData.Email,
				accountData.Pass,
				accountData.Phone,
				accountData.ProviderId,
				accountData.CreatedAt,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.Create(accountData)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO accounts \(id, name, email, password_hash, phone_number, provider_id, created_at\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7\)`).
			WithArgs(
				accountData.ID,
				accountData.Name,
				accountData.Email,
				accountData.Pass,
				accountData.Phone,
				accountData.ProviderId,
				accountData.CreatedAt,
			).WillReturnError(errors.New("insert failed"))

		err = repo.Create(accountData)

		assert.Error(err)
		assert.Equal("insert failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGAccountsRepository_Update(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	accountId := "fake-account-id"
	hashedPass := "hashed-pass"
	phone := "999999999"
	providerId := ""

	accountData := entities.Account{
		ID:               accountId,
		Name:             "jhon",
		Email:            "jhondoe@email.com",
		Pass:             &hashedPass,
		Phone:            &phone,
		ProviderId:       &providerId,
		Is2faEnabled:     false,
		IsEmailVerified:  false,
		LastLoginAt:      nil,
		LastLoginIp:      nil,
		LastLoginCountry: nil,
		LastLoginCity:    nil,
		CreatedAt:        time.Now(),
		UpdatedAt:        nil,
	}

	repo := PGAccountsRepository{Db: db}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`UPDATE accounts SET name = \$1, email = \$2, password_hash = \$3, phone_number = \$4, is_2fa_enabled = \$5, is_email_verified = \$6, last_login_at = \$7, last_login_ip = \$8, last_login_country = \$9, last_login_city = \$10, updated_at = \$11 WHERE id = \$12`).WithArgs(
			accountData.Name,
			accountData.Email,
			accountData.Pass,
			accountData.Phone,
			accountData.Is2faEnabled,
			accountData.IsEmailVerified,
			accountData.LastLoginAt,
			accountData.LastLoginIp,
			accountData.LastLoginCountry,
			accountData.LastLoginCity,
			accountData.UpdatedAt,
			accountData.ID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.Update(accountData)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`UPDATE accounts SET name = \$1, email = \$2, password_hash = \$3, phone_number = \$4, is_2fa_enabled = \$5, is_email_verified = \$6, last_login_at = \$7, last_login_ip = \$8, last_login_country = \$9, last_login_city = \$10, updated_at = \$11 WHERE id = \$12`).
			WithArgs(
				accountData.Name,
				accountData.Email,
				accountData.Pass,
				accountData.Phone,
				accountData.Is2faEnabled,
				accountData.IsEmailVerified,
				accountData.LastLoginAt,
				accountData.LastLoginIp,
				accountData.LastLoginCountry,
				accountData.LastLoginCity,
				accountData.UpdatedAt,
				accountData.ID,
			).WillReturnError(errors.New("update failed"))

		err = repo.Update(accountData)

		assert.Error(err)
		assert.Equal("update failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

package repositories

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/infra/utils"
	"github.com/stretchr/testify/assert"
)

func TestPGVerificationCodesRepository_FindByValue(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGVerificationCodesRepository{Db: db}

	id := "fake-id"
	accountId := "fake-account-id"
	expiresAt := time.Now().Add(15 * time.Minute)
	token, _ := utils.GenerateJWTToken(accountId, expiresAt, "JWT_SECRET")

	data := entities.VerificationCode{
		ID:        id,
		AccountId: accountId,
		Value:     token,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "account_id", "token", "created_at", "expires_at",
		}).AddRow(
			data.ID,
			data.AccountId,
			data.Value,
			data.CreatedAt,
			data.ExpiresAt,
		)

		mock.ExpectQuery(`SELECT id, account_id, token, created_at, expires_at FROM verification_codes WHERE token = \$1`).
			WithArgs(data.Value).WillReturnRows(rows)

		refreshToken := repo.FindByValue(data.Value)

		assert.NotNil(refreshToken)
		assert.Equal(refreshToken.ID, id)
		assert.Equal(refreshToken.Value, token)
		assert.Equal(refreshToken.ExpiresAt, data.ExpiresAt)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, account_id, token, created_at, expires_at FROM verification_codes WHERE token = \$1`).
			WithArgs(data.Value).WillReturnError(errors.New("some error"))

		refreshToken := repo.FindByValue(data.Value)

		assert.Nil(refreshToken)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGVerificationCodesRepository_Create(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGVerificationCodesRepository{Db: db}

	id := "fake-id"
	accountId := "fake-account-id"
	expiresAt := time.Now().Add(15 * time.Minute)
	token, _ := utils.GenerateJWTToken(accountId, expiresAt, "JWT_SECRET")

	data := entities.VerificationCode{
		ID:        id,
		AccountId: accountId,
		Value:     token,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO verification_codes \(id, account_id, token, created_at, expires_at\) VALUES\(\$1, \$2, \$3, \$4, \$5\)`).
			WithArgs(
				data.ID,
				data.AccountId,
				data.Value,
				data.CreatedAt,
				data.ExpiresAt,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Create(data)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO verification_codes \(id, account_id, token, created_at, expires_at\) VALUES\(\$1, \$2, \$3, \$4, \$5\)`).
			WithArgs(
				data.ID,
				data.AccountId,
				data.Value,
				data.CreatedAt,
				data.ExpiresAt,
			).WillReturnError(errors.New("insert failed"))

		err := repo.Create(data)

		assert.Error(err)
		assert.Equal("insert failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGVerificationCodesRepository_Delete(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGVerificationCodesRepository{Db: db}

	id := "fake-id"
	query := `DELETE FROM verification_codes WHERE id = \$1`

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Delete(id)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(id).
			WillReturnError(errors.New("delete failed"))

		err := repo.Delete(id)

		assert.Error(err)
		assert.Equal("delete failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

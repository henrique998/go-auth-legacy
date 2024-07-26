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

func TestPGRefreshTokensRepository_FindByValue(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGRefreshTokensRepository{Db: db}

	id := "fake-id"
	accountId := "fake-account-id"
	expiresAt := time.Now().Add(15 * time.Minute)
	token, _ := utils.GenerateJWTToken(accountId, expiresAt, "JWT_SECRET")

	data := entities.RefreshToken{
		ID:        id,
		AccountId: accountId,
		Value:     token,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "refresh_token", "account_id", "expires_at", "created_at",
		}).AddRow(
			data.ID,
			data.Value,
			data.AccountId,
			data.ExpiresAt,
			data.CreatedAt,
		)

		mock.ExpectQuery(`SELECT id, refresh_token, account_id, expires_at, created_at FROM refresh_tokens WHERE refresh_token = \$1`).
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
		mock.ExpectQuery(`SELECT id, refresh_token, account_id, expires_at, created_at FROM refresh_tokens WHERE refresh_token = \$1`).
			WithArgs(data.Value).WillReturnError(errors.New("some error"))

		refreshToken := repo.FindByValue(data.Value)

		assert.Nil(refreshToken)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGRefreshTokensRepository_Create(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGRefreshTokensRepository{Db: db}

	id := "fake-id"
	accountId := "fake-account-id"
	expiresAt := time.Now().Add(15 * time.Minute)
	token, _ := utils.GenerateJWTToken(accountId, expiresAt, "JWT_SECRET")

	data := entities.RefreshToken{
		ID:        id,
		AccountId: accountId,
		Value:     token,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO refresh_tokens \(id, refresh_token, account_id, expires_at, created_at\) VALUES\(\$1, \$2, \$3, \$4, \$5\)`).
			WithArgs(
				data.ID,
				data.Value,
				data.AccountId,
				data.ExpiresAt,
				data.CreatedAt,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Create(data)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO refresh_tokens \(id, refresh_token, account_id, expires_at, created_at\) VALUES\(\$1, \$2, \$3, \$4, \$5\)`).
			WithArgs(
				data.ID,
				data.Value,
				data.AccountId,
				data.ExpiresAt,
				data.CreatedAt,
			).WillReturnError(errors.New("insert failed"))

		err := repo.Create(data)

		assert.Error(err)
		assert.Equal("insert failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGRefreshTokensRepository_Delete(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGRefreshTokensRepository{Db: db}

	val := "fake-val"
	query := `DELETE FROM refresh_tokens WHERE refresh_token = \$1`

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(val).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Delete(val)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(query).
			WithArgs(val).
			WillReturnError(errors.New("delete failed"))

		err := repo.Delete(val)

		assert.Error(err)
		assert.Equal("delete failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

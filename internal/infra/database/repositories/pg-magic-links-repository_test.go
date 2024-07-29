package repositories

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/infra/utils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestPGMagicLinksRepository_FindByValue(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGMagicLinksRepository{Db: db}

	id := "fake-magic-link-id"
	code, _ := utils.GenerateCode(10)
	expiresAt := time.Now().Add(15 * time.Minute)

	data := entities.MagicLink{
		ID:        id,
		AccountId: "fake-account-id",
		Code:      code,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "account_id", "code", "expires_at", "created_at",
		}).AddRow(
			data.ID,
			data.AccountId,
			data.Code,
			data.ExpiresAt,
			data.CreatedAt,
		)

		mock.ExpectQuery(`SELECT id, account_id, code, expires_at, created_at FROM magic_links WHERE code = \$1`).
			WithArgs(data.Code).WillReturnRows(rows)

		link := repo.FindByValue(data.Code)

		assert.NotNil(link)
		assert.Equal(link.ID, id)
		assert.Equal(link.Code, code)
		assert.Equal(link.ExpiresAt, data.ExpiresAt)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, account_id, code, expires_at, created_at FROM magic_links WHERE code = \$1`).
			WithArgs(data.Code).WillReturnError(errors.New("some error"))

		link := repo.FindByValue(data.Code)

		assert.Nil(link)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGMagicLinksRepository_Create(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGMagicLinksRepository{Db: db}

	code, _ := utils.GenerateCode(10)

	data := entities.MagicLink{
		ID:        "fake-magic-link-id",
		AccountId: "fake-account-id",
		Code:      code,
		ExpiresAt: time.Now().Add(15 * time.Minute),
		CreatedAt: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO magic_links \(id, account_id, code, expires_at, created_at\) VALUES \(\$1, \$2, \$3, \$4, \$5\)`).
			WithArgs(
				data.ID,
				data.AccountId,
				data.Code,
				data.ExpiresAt,
				data.CreatedAt,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Create(data)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO magic_links \(id, account_id, code, expires_at, created_at\) VALUES \(\$1, \$2, \$3, \$4, \$5\)`).
			WithArgs(
				data.ID,
				data.AccountId,
				data.Code,
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

func TestPGMagicLinksRepository_Delete(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGMagicLinksRepository{Db: db}

	id := "fake-magic-link-id"

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`DELETE FROM magic_links WHERE id = \$1`).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Delete(id)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`DELETE FROM magic_links WHERE id = \$1`).
			WithArgs(id).
			WillReturnError(errors.New("delete failed"))

		err := repo.Delete(id)

		assert.Error(err)
		assert.Equal("delete failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

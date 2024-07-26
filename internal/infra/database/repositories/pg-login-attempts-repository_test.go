package repositories

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/stretchr/testify/assert"
)

func TestPGLoginAttemptsRepository_Create(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGLoginAttemptsRepository{Db: db}

	data := entities.LoginAttempt{
		ID:          "fake-attempt-id",
		Email:       "jhondoe@email.com",
		IPAddress:   "fake-attempt-ip",
		UserAgent:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		Success:     true,
		AttemptedAt: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO login_attempts \(id, email, ip_address, user_agent, success, attempted_at\)
		VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\)`).
			WithArgs(
				data.ID,
				data.Email,
				data.IPAddress,
				data.UserAgent,
				data.Success,
				data.AttemptedAt,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.Create(data)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO login_attempts \(id, email, ip_address, user_agent, success, attempted_at\)
		VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\)`).
			WithArgs(
				data.ID,
				data.Email,
				data.IPAddress,
				data.UserAgent,
				data.Success,
				data.AttemptedAt,
			).WillReturnError(errors.New("insert failed"))

		err = repo.Create(data)

		assert.Error(err)
		assert.Equal("insert failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

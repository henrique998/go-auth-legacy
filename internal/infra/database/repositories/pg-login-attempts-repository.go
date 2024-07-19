package repositories

import (
	"database/sql"

	"github.com/henrique998/go-auth/internal/app/entities"
)

type PGLoginAttemptsRepository struct {
	Db *sql.DB
}

func (r *PGLoginAttemptsRepository) Create(la entities.LoginAttempt) error {
	query :=
		`INSERT INTO login_attempts (id, email, ip_address, user_agent, success, attempted_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.Db.Exec(query,
		la.ID,
		la.Email,
		la.IPAddress,
		la.UserAgent,
		la.Success,
		la.AttemptedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

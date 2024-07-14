package repositories

import (
	"database/sql"

	"github.com/henrique998/go-auth/internal/app/entities"
)

type PGMagicLinksRepository struct {
	Db *sql.DB
}

func (r *PGMagicLinksRepository) Create(ml entities.MagicLink) error {
	query :=
		`INSERT INTO magic_links (id, account_id, code, expires_at, created_at) VALUES ($1, $2, $3, $4, $5)`

	_, err := r.Db.Exec(query,
		ml.ID,
		ml.AccountId,
		ml.Code,
		ml.ExpiresAt,
		ml.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

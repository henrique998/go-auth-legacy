package repositories

import (
	"database/sql"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type PGMagicLinksRepository struct {
	Db *sql.DB
}

func (r *PGMagicLinksRepository) FindByValue(val string) *entities.MagicLink {
	var magicLink entities.MagicLink

	query := "SELECT id, account_id, code, expires_at, created_at FROM magic_links WHERE code = $1"
	row := r.Db.QueryRow(query, val)

	err := row.Scan(&magicLink.ID, &magicLink.AccountId, &magicLink.Code, &magicLink.ExpiresAt, &magicLink.CreatedAt)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error("Error trying to find refresh token", err)
		}
		return nil
	}

	return &magicLink
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

func (r *PGMagicLinksRepository) Delete(id string) error {
	query := "DELETE FROM magic_links WHERE id = $1"

	_, err := r.Db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

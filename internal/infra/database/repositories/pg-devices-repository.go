package repositories

import (
	"database/sql"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type PGDevicesRepository struct {
	Db *sql.DB
}

func (r *PGDevicesRepository) FindByIpAndAccountId(ip, accountId string) (*entities.Device, errors.IAppError) {
	var device entities.Device

	query := "SELECT * FROM devices WHERE ip_address = $1 AND account_id = $2 LIMIT 1"

	row := r.Db.QueryRow(query, ip, accountId)

	err := row.Scan(
		&device.ID,
		&device.AccountID,
		&device.DeviceName,
		&device.UserAgent,
		&device.Platform,
		&device.IPAddress,
		&device.CreatedAt,
		&device.UpdatedAt,
		&device.LastLoginAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		logger.Error("Error trying to find account!", err)
		return nil, errors.NewAppError(err.Error(), 500)
	}

	return &device, nil
}

func (r *PGDevicesRepository) Create(device entities.Device) errors.IAppError {
	query :=
		`INSERT INTO devices (id, account_id, device_name, user_agent, platform, ip_address, created_at, updated_at, last_login_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := r.Db.Exec(query,
		device.ID,
		device.AccountID,
		device.DeviceName,
		device.UserAgent,
		device.Platform,
		device.IPAddress,
		device.CreatedAt,
		*device.UpdatedAt,
		device.LastLoginAt,
	)
	if err != nil {
		return errors.NewAppError(err.Error(), 500)
	}

	return nil
}

func (r *PGDevicesRepository) Update(device entities.Device) errors.IAppError {
	query := "UPDATE devices SET device_name = $1, updated_at = $2 WHERE id = $3"

	_, err := r.Db.Exec(
		query,
		device.DeviceName,
		device.UpdatedAt,
		device.ID,
	)
	if err != nil {
		return errors.NewAppError(err.Error(), 500)
	}

	return nil
}

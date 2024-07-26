package repositories

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/stretchr/testify/assert"
)

func TestPGDevicesRepository_FindByIpAndAccountId(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	deviceId := "fake-device-id"
	accountId := "fake-account-id"
	ip := "fake-device-ip"

	deviceData := entities.Device{
		ID:          deviceId,
		AccountID:   accountId,
		DeviceName:  "device-name",
		UserAgent:   "device-user-agent",
		Platform:    "device-platform",
		IPAddress:   ip,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
		LastLoginAt: time.Now(),
	}

	repo := PGDevicesRepository{
		Db: db,
	}

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "account_id", "device_name", "user_agent", "platform", "ip_address", "created_at", "updated_at", "last_login_at",
		}).AddRow(
			deviceData.ID,
			deviceData.AccountID,
			deviceData.DeviceName,
			deviceData.UserAgent,
			deviceData.Platform,
			deviceData.IPAddress,
			deviceData.CreatedAt,
			deviceData.UpdatedAt,
			deviceData.LastLoginAt,
		)

		mock.ExpectQuery(`SELECT \* FROM devices WHERE ip_address = \$1 AND account_id = \$2 LIMIT 1`).
			WithArgs(ip, accountId).
			WillReturnRows(rows)

		device := repo.FindByIpAndAccountId(ip, accountId)

		assert.NotNil(device)
		assert.Equal(accountId, device.AccountID)
		assert.Equal(ip, device.IPAddress)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM devices WHERE ip_address = \$1 AND account_id = \$2 LIMIT 1`).
			WithArgs(ip, accountId).
			WillReturnError(errors.New("some other error"))

		device := repo.FindByIpAndAccountId(ip, accountId)

		assert.Nil(device)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGDevicesRepository_FindManyByAccountId(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	accountId := "fake-account-id"

	deviceData := entities.Device{
		ID:          "fake-device-id",
		AccountID:   accountId,
		DeviceName:  "device-name",
		UserAgent:   "device-user-agent",
		Platform:    "device-platform",
		IPAddress:   "fake-device-ip",
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
		LastLoginAt: time.Now(),
	}

	repo := PGDevicesRepository{
		Db: db,
	}

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "account_id", "device_name", "user_agent", "platform", "ip_address", "created_at", "updated_at", "last_login_at",
		}).AddRow(
			deviceData.ID,
			deviceData.AccountID,
			deviceData.DeviceName,
			deviceData.UserAgent,
			deviceData.Platform,
			deviceData.IPAddress,
			deviceData.CreatedAt,
			deviceData.UpdatedAt,
			deviceData.LastLoginAt,
		)

		mock.ExpectQuery(`SELECT id, account_id, device_name, user_agent, platform, ip_address, created_at, updated_at, last_login_at FROM devices WHERE account_id = \$1`).
			WithArgs(accountId).
			WillReturnRows(rows)

		devices := repo.FindManyByAccountId(accountId)

		assert.Len(devices, 1)
		assert.Equal(devices[0].ID, "fake-device-id")

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGDevicesRepository_Create(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGDevicesRepository{Db: db}

	deviceData := entities.Device{
		ID:          "fake-device-id",
		AccountID:   "fake-account-id",
		DeviceName:  "device-name",
		UserAgent:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		Platform:    "Windows",
		IPAddress:   "fake-device-ip",
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
		LastLoginAt: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO devices \(id, account_id, device_name, user_agent, platform, ip_address, created_at, updated_at, last_login_at\)
		VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9\)`).
			WithArgs(
				deviceData.ID,
				deviceData.AccountID,
				deviceData.DeviceName,
				deviceData.UserAgent,
				deviceData.Platform,
				deviceData.IPAddress,
				deviceData.CreatedAt,
				deviceData.UpdatedAt,
				deviceData.LastLoginAt,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.Create(deviceData)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO devices \(id, account_id, device_name, user_agent, platform, ip_address, created_at, updated_at, last_login_at\)
		VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9\)`).
			WithArgs(
				deviceData.ID,
				deviceData.AccountID,
				deviceData.DeviceName,
				deviceData.UserAgent,
				deviceData.Platform,
				deviceData.IPAddress,
				deviceData.CreatedAt,
				deviceData.UpdatedAt,
				deviceData.LastLoginAt,
			).WillReturnError(errors.New("insert failed"))

		err = repo.Create(deviceData)

		assert.Error(err)
		assert.Equal("insert failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGDevicesRepository_Update(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGDevicesRepository{Db: db}

	deviceData := entities.Device{
		ID:          "fake-device-id",
		AccountID:   "fake-account-id",
		DeviceName:  "updated-device-name",
		UserAgent:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		Platform:    "Windows",
		IPAddress:   "fake-device-ip",
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
		LastLoginAt: time.Now(),
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`UPDATE devices SET device_name = \$1, updated_at = \$2 WHERE id = \$3`).
			WithArgs(
				deviceData.DeviceName,
				deviceData.UpdatedAt,
				deviceData.ID,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.Update(deviceData)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`UPDATE devices SET device_name = \$1, updated_at = \$2 WHERE id = \$3`).
			WithArgs(
				deviceData.DeviceName,
				deviceData.UpdatedAt,
				deviceData.ID,
			).WillReturnError(errors.New("update failed"))

		err = repo.Update(deviceData)

		assert.Error(err)
		assert.Equal("update failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

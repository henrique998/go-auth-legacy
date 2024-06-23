package contracts

import (
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/errors"
)

type DevicesRepository interface {
	FindByIpAndAccountId(ip, accountId string) (*entities.Device, errors.IAppError)
	Create(device entities.Device) errors.IAppError
	Update(device entities.Device) errors.IAppError
}

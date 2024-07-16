package contracts

import (
	"github.com/henrique998/go-auth/internal/app/entities"
)

type DevicesRepository interface {
	FindByIpAndAccountId(ip, accountId string) *entities.Device
	FindManyByAccountId(accountId string) []entities.Device
	Create(device entities.Device) error
	Update(device entities.Device) error
}

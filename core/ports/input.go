package ports

import "github.com/awjmj/address-api/core/domain"

type GetAddressInputPort interface {
	Get(zipCode string) (*domain.Address, error)
}

package mapper

import (
	"github.com/awjmj/address-api/core/domain"
	"github.com/awjmj/address-api/infra/client/response"
)

func ToAddress(address *response.AddressResponse) *domain.Address {

	return &domain.Address{
		Street:        address.Street,
		Neighboorhood: address.Neighboorhood,
		City:          address.City,
		State:         address.State,
		Complement:    address.Complement,
		ZipCode:       address.ZipCode,
	}
}

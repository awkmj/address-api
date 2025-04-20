package usecase

import (
	"log/slog"

	adapter "github.com/awjmj/address-api/adapter/api"
	"github.com/awjmj/address-api/core/domain"
	"github.com/awjmj/address-api/core/ports"
)

type GetAddressUseCase struct {
	getAddress ports.GetAddressInputPort
}

func NewGetAddressUseCase() *GetAddressUseCase {

	getAddress := adapter.NewGetAddress()
	return &GetAddressUseCase{
		getAddress: getAddress,
	}
}

func (g *GetAddressUseCase) Get(zipCode string) (*domain.Address, error) {

	slog.Info("[GET ADDRESS USE CASE][START]", "zip_code", zipCode)

	address, err := g.getAddress.Get(zipCode)

	if err != nil {
		slog.Info("[GET ADDRESS USE CASE][ERROR]", "error", err.Error())
		return nil, err
	}

	slog.Info("[GET ADDRESS USE CASE][END]", "zip_code", zipCode)
	return address, nil
}

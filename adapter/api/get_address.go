package adapter

import (
	"encoding/json"
	"io"
	"log/slog"

	"github.com/awjmj/address-api/adapter/mapper"
	"github.com/awjmj/address-api/core/domain"
	"github.com/awjmj/address-api/infra/client"
	"github.com/awjmj/address-api/infra/client/response"
)

type GetAddress struct {
	client *client.HttpClient
}

func NewGetAddress() *GetAddress {

	url := "https://viacep.com.br"
	httpClient := client.NewHttpClient(client.WithUrl(url))
	return &GetAddress{client: httpClient}
}

func (g *GetAddress) Get(zipCode string) (*domain.Address, error) {

	slog.Info("[GET ADDRESS ADAPTER][START]", "zip_code", zipCode)

	res, err := g.client.Get(zipCode)

	if err != nil {
		slog.Info("[GET ADDRESS ADAPTER][REQUEST ERROR]", "error", err.Error())
		return nil, err
	}

	result, err := io.ReadAll(res)

	if err != nil {
		slog.Info("[GET ADDRESS ADAPTER][PARSE ERROR]", "error", err.Error())
		return nil, err
	}

	address := &response.AddressResponse{}
	json.Unmarshal(result, &address)
	slog.Info("[GET ADDRESS ADAPTER][END]", "zip_code", zipCode)
	return mapper.ToAddress(address), nil
}

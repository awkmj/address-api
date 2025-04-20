package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/awjmj/address-api/usecase"
)

type AddressController struct {
	getAddressUseCase *usecase.GetAddressUseCase
}

func NewAddressController() *AddressController {

	return &AddressController{
		getAddressUseCase: usecase.NewGetAddressUseCase(),
	}
}

func (a *AddressController) Get(w http.ResponseWriter, r *http.Request) {

	zipCode := r.URL.Query().Get("zipCode")
	address, err := a.getAddressUseCase.Get(zipCode)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error": %s}`, err.Error())))
		return
	}

	data, _ := json.Marshal(address)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

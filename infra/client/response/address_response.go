package response

type AddressResponse struct {
	Street        string `json:"logradouro"`
	Neighboorhood string `json:"bairro"`
	City          string `json:"localidade"`
	State         string `json:"uf"`
	Complement    string `json:"complemento"`
	ZipCode       string `json:"cep"`
}

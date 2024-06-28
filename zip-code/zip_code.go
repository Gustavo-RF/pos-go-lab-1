package zipcode

import (
	"github.com/Gustavo-RF/pos-go-lab-1/internal"
	"github.com/Gustavo-RF/pos-go-lab-1/zip-code/entities"
)

type ZipCodeResponse struct {
	Localidade string `json:"localidade"`
}

func GetZipCode(zipcode string) (*ZipCodeResponse, error) {

	zipCodeApiResponse, err := fetch(zipcode)

	if err != nil {
		return nil, err
	}

	response := ZipCodeResponse{
		Localidade: zipCodeApiResponse.Localidade,
	}

	return &response, nil
}

func fetch(zipcode string) (*entities.ZipCodeApiResponse, error) {
	res, err := internal.Request("https://viacep.com.br/ws/"+zipcode+"/json/", "GET")

	if err != nil {
		return nil, err
	}

	zipCodeApiResponse, err := entities.NewZipCodeApiResponse(res)

	if err != nil {
		return nil, err
	}

	return zipCodeApiResponse, nil
}

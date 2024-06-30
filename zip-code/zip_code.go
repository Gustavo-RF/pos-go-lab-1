package zipcode

import (
	"errors"

	"github.com/Gustavo-RF/pos-go-lab-1/internal/web"
	"github.com/Gustavo-RF/pos-go-lab-1/zip-code/entities"
)

func GetZipCode(zipcode string) (*entities.ZipCodeResponse, error) {

	zipCodeApiResponse, err := fetch(zipcode)

	if err != nil {
		return nil, err
	}

	response := entities.NewZipCodeResponse(zipCodeApiResponse.Localidade)

	return &response, nil
}

func fetch(zipcode string) (*entities.ZipCodeApiResponse, error) {
	res, err := web.Request("https://viacep.com.br/ws/"+zipcode+"/json/", "GET")

	if err != nil {
		return nil, err
	}

	zipCodeApiResponse, err := entities.NewZipCodeApiResponse(res)
	if err != nil {
		return nil, err
	}

	if zipCodeApiResponse.Erro == "true" {
		return nil, errors.New("zipcode not found")
	}

	return zipCodeApiResponse, nil
}

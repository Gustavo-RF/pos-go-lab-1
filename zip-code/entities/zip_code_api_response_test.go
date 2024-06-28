package entities

import (
	"testing"

	"github.com/Gustavo-RF/pos-go-lab-1/internal"
	"github.com/stretchr/testify/assert"
)

func TestRequestShouldReturnValidField(t *testing.T) {
	res, err := internal.Request("https://viacep.com.br/ws/29092260/json/", "GET")

	assert.Nil(t, err)

	zipCodeApiResponse, err := NewZipCodeApiResponse(res)

	assert.Nil(t, err)

	assert.Equal(t, zipCodeApiResponse.Cep, "29092-260")
	assert.Equal(t, zipCodeApiResponse.Localidade, "Vitória")
}

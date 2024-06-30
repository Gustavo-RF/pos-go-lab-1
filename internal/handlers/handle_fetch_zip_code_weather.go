package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Gustavo-RF/pos-go-lab-1/weather"
	zipcode "github.com/Gustavo-RF/pos-go-lab-1/zip-code"
	"github.com/paemuri/brdoc"
)

type Response struct {
	Message string `json:"message"`
}

func HandleFetchZipCodeTemp(res http.ResponseWriter, req *http.Request, weatherApiKey string) {
	cep := req.URL.Query().Get("cep")

	if cep == "" {
		res.WriteHeader(http.StatusUnprocessableEntity)
		response := Response{
			Message: "Cep is required",
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	if len(cep) != 8 {
		res.WriteHeader(http.StatusUnprocessableEntity)
		response := Response{
			Message: "Invalid zipcode",
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	if !brdoc.IsCEP(cep) {
		res.WriteHeader(http.StatusUnprocessableEntity)
		response := Response{
			Message: "Invalid zipcode",
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	cepFind, err := zipcode.GetZipCode(cep)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		response := Response{
			Message: err.Error(),
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	weather, err := weather.GetWeather(cepFind.Localidade, weatherApiKey)

	if err != nil {
		res.WriteHeader(http.StatusBadGateway)
		response := Response{
			Message: "Error while get weather: " + err.Error(),
		}
		json.NewEncoder(res).Encode(response)
		return
	}

	res.Header().Set("Content-type", "application/json")
	json.NewEncoder(res).Encode(weather)
}

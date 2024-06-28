package weather

import (
	"github.com/Gustavo-RF/pos-go-lab-1/internal"
	"github.com/Gustavo-RF/pos-go-lab-1/weather/entities"
)

type WeatherResponse struct {
	TempC float32 `json:"temp_c"`
	TempF float32 `json:"temp_f"`
}

func GetWeather(local string) (*WeatherResponse, error) {
	key := "teste"
	weatherApiResponse, err := fetch(key, local)

	if err != nil {
		return nil, err
	}

	wheaterResponse := WeatherResponse{
		TempC: weatherApiResponse.Current.TempC,
		TempF: weatherApiResponse.Current.TempC,
	}

	return &wheaterResponse, nil
}

func fetch(key, local string) (*entities.WeatherApiResponse, error) {
	res, err := internal.Request("http://api.weatherapi.com/v1?key="+key+"&q="+local, "GET")

	if err != nil {
		return nil, err
	}

	weatherApiResponse, err := entities.NewWeatherApiResponse(res)

	if err != nil {
		return nil, err
	}

	return weatherApiResponse, nil
}

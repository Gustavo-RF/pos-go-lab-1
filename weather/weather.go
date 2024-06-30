package weather

import (
	"fmt"
	"net/url"

	"github.com/Gustavo-RF/pos-go-lab-1/internal/web"
	"github.com/Gustavo-RF/pos-go-lab-1/weather/entities"
)

func GetWeather(local, key string) (entities.WeatherResponse, error) {
	weatherApiResponse, err := fetch(key, local)

	if err != nil {
		return entities.WeatherResponse{}, err
	}

	wheaterResponse := entities.NewWeatherResponse(weatherApiResponse.Current.TempC)

	return wheaterResponse, nil
}

func fetch(key, local string) (*entities.WeatherApiResponse, error) {
	localEscaped := url.QueryEscape(local)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", key, localEscaped)

	res, err := web.Request(url, "GET")
	if err != nil {
		return nil, err
	}
	weatherApiResponse, err := entities.NewWeatherApiResponse(res)

	if err != nil {
		return nil, err
	}

	return weatherApiResponse, nil
}

package weather

import (
	"fmt"
	"net/url"

	"github.com/Gustavo-RF/pos-go-lab-1/weather/entities"
)

type RequestFunc func(url, method string) ([]byte, error)

func GetWeather(local, key string, requestFunc RequestFunc) (entities.WeatherResponse, error) {
	weatherApiResponse, err := fetch(key, local, requestFunc)
	if err != nil {
		return entities.WeatherResponse{}, err
	}

	weatherResponse := entities.NewWeatherResponse(weatherApiResponse.Current.TempC)
	return weatherResponse, nil
}

func fetch(key, local string, requestFunc RequestFunc) (*entities.WeatherApiResponse, error) {
	localEscaped := url.QueryEscape(local)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", key, localEscaped)

	res, err := requestFunc(url, "GET")
	if err != nil {
		return nil, err
	}

	weatherApiResponse, err := entities.NewWeatherApiResponse(res)
	if err != nil {
		return nil, err
	}

	return weatherApiResponse, nil
}

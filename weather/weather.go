package weather

import (
	"fmt"
	"net/url"

	"github.com/Gustavo-RF/pos-go-lab-1/weather/entities"
)

type RequestFunc func(url, method string) ([]byte, error)

func GetWeather(local string, requestFunc RequestFunc) (entities.WeatherResponse, error) {
	weatherApiResponse, err := fetch(local, requestFunc)
	if err != nil {
		return entities.WeatherResponse{}, err
	}

	weatherResponse := entities.NewWeatherResponse(weatherApiResponse.Current.TempC)
	return weatherResponse, nil
}

func fetch(local string, requestFunc RequestFunc) (*entities.WeatherApiResponse, error) {
	localEscaped := url.QueryEscape(local)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=2d29c8f65271404488800806242506&q=%s", localEscaped)

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

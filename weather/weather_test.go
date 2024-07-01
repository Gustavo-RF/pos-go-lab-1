package weather

import (
	"errors"
	"testing"

	"github.com/Gustavo-RF/pos-go-lab-1/internal/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetWeather(t *testing.T) {
	mockRequestFunc := new(web.MockRequestFunc)
	mockRequestFunc.On("Request", mock.Anything, "GET").Return([]byte(`{"current": {"temp_c": 25.0}}`), nil)

	resp, err := GetWeather("London", mockRequestFunc.Request)
	assert.NoError(t, err)
	assert.Equal(t, float32(25.0), resp.TempC)

	mockRequestFunc.AssertExpectations(t)
}

func TestGetWeather_RequestError(t *testing.T) {
	mockRequestFunc := new(web.MockRequestFunc)
	mockRequestFunc.On("Request", mock.Anything, "GET").Return([]byte{}, errors.New("request error"))

	_, err := GetWeather("London", mockRequestFunc.Request)
	assert.Error(t, err)
	assert.Equal(t, "request error", err.Error())

	mockRequestFunc.AssertExpectations(t)
}

func TestFetch(t *testing.T) {
	mockRequestFunc := new(web.MockRequestFunc)
	mockRequestFunc.On("Request", mock.Anything, "GET").Return([]byte(`{"current": {"temp_c": 25.0}}`), nil)

	resp, err := fetch("London", mockRequestFunc.Request)
	assert.NoError(t, err)
	assert.Equal(t, float32(25.0), resp.Current.TempC)

	mockRequestFunc.AssertExpectations(t)
}

func TestFetch_RequestError(t *testing.T) {
	mockRequestFunc := new(web.MockRequestFunc)
	mockRequestFunc.On("Request", mock.Anything, "GET").Return([]byte{}, errors.New("request error"))

	_, err := fetch("London", mockRequestFunc.Request)
	assert.Error(t, err)
	assert.Equal(t, "request error", err.Error())

	mockRequestFunc.AssertExpectations(t)
}

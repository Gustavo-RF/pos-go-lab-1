package web

import "github.com/stretchr/testify/mock"

type MockRequestFunc struct {
	mock.Mock
}

func (m *MockRequestFunc) Request(url, method string) ([]byte, error) {
	args := m.Called(url, method)
	return args.Get(0).([]byte), args.Error(1)
}

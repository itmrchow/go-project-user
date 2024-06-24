package handler

import "github.com/stretchr/testify/mock"

type EncryptionHandlerMock struct {
	mock.Mock
}

func NewEncryptionHandlerMock() *EncryptionHandlerMock {
	return &EncryptionHandlerMock{}
}

func (e *EncryptionHandlerMock) HashPassword(password string) (string, error) {
	args := e.Called(password)
	return args.String(0), args.Error(1)
}

func (e *EncryptionHandlerMock) CheckPasswordHash(password string, hash string) bool {
	args := e.Called(password, hash)
	return args.Bool(0)
}

package usecase

import (
	"errors"
)

type UseCaseError error

var (
	ErrUserAlreadyExists = errors.New("user already exists")

	ErrUserNotExists = errors.New("user not exists")

	ErrUnauthorized = errors.New("unauthorized")

	ErrDbFail = errors.New("db fail")

	ErrDbInsertFail = errors.New("db fail, insert key exist")

	ErrPasswordHash = errors.New("password hash fail")
)

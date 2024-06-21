package usecase

import (
	"errors"

	"itmrchow/go-project/user/src/usecase/repo"
)

type PingService interface {
	Ping() (string, error)
}

type PingServiceImpl struct {
	userRepo repo.UserRepo
}

func NewPingServiceImpl(userRepo repo.UserRepo) *PingServiceImpl {
	return &PingServiceImpl{userRepo: userRepo}
}

func (p *PingServiceImpl) Ping() (string, error) {
	_, err := p.userRepo.Get("for-chec")

	if err != nil {
		return "error", errors.Join(ErrDbFail, err)
	}

	return "pong", nil
}

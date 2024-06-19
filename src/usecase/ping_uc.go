package usecase

import "errors"

type PingService interface {
	Ping() (string, error)
}

type PingServiceImpl struct {
}

func (p *PingServiceImpl) Ping() (string, error) {
	return "pong", errors.New("new jeff error")
}

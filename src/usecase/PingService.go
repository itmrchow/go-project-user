package usecase

type PingService interface {
	Ping() string
}

type PingServiceImpl struct {
}

func (p *PingServiceImpl) Ping() string {
	return "pong"
}

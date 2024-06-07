package controllers

import "itmrchow/go-project/user/src/usecase"

type PingController struct {
	pingService usecase.PingService
}

func NewPingController() *PingController {
	service := &usecase.PingServiceImpl{}
	return &PingController{service}
}

func (controller *PingController) Ping() string {
	return controller.pingService.Ping()
}

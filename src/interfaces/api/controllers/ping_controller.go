package controllers

import "itmrchow/go-project/user/src/usecase"

type PingController struct {
	pingService usecase.PingService
}

func NewPingController(pingService usecase.PingService) *PingController {
	return &PingController{pingService: pingService}
}

func (controller *PingController) Ping() (string, error) {
	return controller.pingService.Ping()
}

func (controller *PingController) GetHelloWorld() string {
	return "HelloWorld"
}

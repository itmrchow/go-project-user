package api

import (
	"github.com/gin-gonic/gin"
	// "itmrchow/go-project/user/src/interfaces/controllers"
)

var router = gin.Default()

func Run() {

	getRoutes()

	router.Run(":8080")
}

func getRoutes() {
	apiV1 := router.Group("api/v1/")
	// controllers := controllers.NewUserController()

	// ping
	addPingRoutes(apiV1)
	// user
	addUserRoutes(apiV1)
}

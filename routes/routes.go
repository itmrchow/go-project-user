package routes

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Run() {

	getRoutes()

	router.Run(":8080")
}

func getRoutes() {
	apiV1 := router.Group("api/v1/")

	// ping
	addPingRoutes(apiV1)
	// user
	addUserRoutes(apiV1)

}

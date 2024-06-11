package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"itmrchow/go-project/user/docs"

)

// @title           User service API
// @version         1.0
// @description     User service

// @host      localhost:8080

// @securityDefinitions.basic  BasicAuth

var router = gin.Default()

func Run() {

	getRoutes()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}

func getRoutes() {
	apiV1 := router.Group("api/v1/")

	docs.SwaggerInfo.BasePath = apiV1.BasePath()

	// Example
	addExampleRoutes(apiV1)

	// user
	addUserRoutes(apiV1)

}

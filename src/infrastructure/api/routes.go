package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"itmrchow/go-project/user/docs"
)

var router = gin.Default()

func Run() {

	router.Use(ErrorHandle())

	getRoutes()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}

func getRoutes() {
	apiV1 := router.Group("api/v1/")

	docs.SwaggerInfo.BasePath = apiV1.BasePath()

	//
	// Example
	addExampleRoutes(apiV1)

	// user
	addUserRoutes(apiV1)

}

package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"itmrchow/go-project/user/docs"
)

var router = gin.Default()

func Run() {

	getRoutes()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func getRoutes() {
	apiV1 := router.Group("api/v1/")
	docs.SwaggerInfo.BasePath = apiV1.BasePath()
	// controllers := controllers.NewUserController()

	// ping
	addPingRoutes(apiV1)
	// user
	addUserRoutes(apiV1)

}

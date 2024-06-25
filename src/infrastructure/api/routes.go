package api

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"itmrchow/go-project/user/docs"
	"itmrchow/go-project/user/src/infrastructure/api/reqdto"
	"itmrchow/go-project/user/src/usecase"
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

func GetAuthUser(c *gin.Context) (*reqdto.AuthUser, error) {
	authUserInfo, isExists := c.Get("AuthUser")
	if !isExists {
		return &reqdto.AuthUser{}, usecase.ErrUnauthorized
	}

	authUser, ok := authUserInfo.(reqdto.AuthUser)
	if !ok {
		return &reqdto.AuthUser{}, usecase.ErrUnauthorized
	}

	return &authUser, nil
}
